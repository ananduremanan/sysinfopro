<script lang="ts">
	import { bytesToGb } from '$lib';
	import Modal from '../ui/Modal.svelte';

	let showConfirmation = $state(false);

	interface Files {
		[key: string]: any[];
	}

	interface FileData {
		Files: Files;
		TotalSize: number;
	}

	interface PieProps {
		files: FileData;
	}

	let { files = { Files: { '': [] }, TotalSize: 0 } }: PieProps = $props();

	// Calculate sizes per category
	let chartData: any = $state([]);
	let colors = $state(['#4F46E5', '#10B981', '#F59E0B', '#EF4444', '#8B5CF6', '#EC4899']);

	function prepareChartData() {
		chartData = [];

		// Process file data for chart
		if (files && files.Files) {
			Object.entries(files.Files).forEach(([key, value], index) => {
				const itemCount = value.length;
				const totalItems = Object.values(files.Files).reduce((sum, arr) => sum + arr.length, 0);
				const estimatedSize = files.TotalSize * (itemCount / totalItems);

				chartData.push({
					name: key,
					value: estimatedSize,
					color: colors[index % colors.length],
					count: itemCount
				});
			});
		}
	}

	// Initialize chart on mount
	$effect(() => {
		prepareChartData();
		drawChart();
	});

	function drawChart() {
		const canvas: any = document.getElementById('pieChart');
		if (!canvas) return;

		const ctx = canvas.getContext('2d');
		const centerX = canvas.width / 2;
		const centerY = canvas.height / 2;
		const radius = Math.min(centerX, centerY) - 10;

		let total = chartData.reduce((sum: any, item: any) => sum + item.value, 0);
		let startAngle = 0;

		ctx.clearRect(0, 0, canvas.width, canvas.height);

		// Draw pie slices
		chartData.forEach((item: any) => {
			const sliceAngle = (2 * Math.PI * item.value) / total;

			ctx.beginPath();
			ctx.moveTo(centerX, centerY);
			ctx.arc(centerX, centerY, radius, startAngle, startAngle + sliceAngle);
			ctx.closePath();

			ctx.fillStyle = item.color;
			ctx.fill();

			startAngle += sliceAngle;
		});
	}
</script>

<div class="grid grid-cols-2">
	<div class="relative w-full flex justify-center items-center">
		<canvas id="pieChart" width="100" height="100" class="mx-auto"></canvas>
	</div>

	<div class="mt-2 w-full">
		<div class="flex flex-col text-xs gap-1 justify-between">
			{#each chartData.slice(-3) as item}
				<div class="flex items-center">
					<div class="w-2 h-2 rounded-full mr-2" style="background-color: {item.color}"></div>
					<span class="flex-1">{item.name}</span>
					<span class="ml-4 font-medium">{bytesToGb(item.value)} GB</span>
				</div>
			{/each}
			<div class="flex flex-col justify-between mt-2 border-gray-200 font-bold self-end">
				<span>Total {bytesToGb(files.TotalSize)} GB</span>
				<button
					class="bg-blue-500 px-2 py-1 rounded-2xl mt-3 text-white text-xs cursor-pointer"
					onclick={() => (showConfirmation = !showConfirmation)}
				>
					Free Space</button
				>
			</div>
		</div>
	</div>

	<Modal showModal={showConfirmation} showClose={false}>
		<div>
			<div class="text-sm font-bold mb-2">Warning!</div>
			<div>This Will Clear All Temp and Browser Cache and not reversable</div>
			<div class="w-full flex justify-center items-center gap-2 mt-2">
				<button
					class="px-2 py-1 bg-red-500 rounded-lg text-white cursor-pointer"
					onclick={() => (showConfirmation = !showConfirmation)}>Clear Space</button
				>
				<button
					class="px-2 py-1 border border-red-500 rounded-lg cursor-pointer"
					onclick={() => (showConfirmation = !showConfirmation)}>Cancel</button
				>
			</div>
		</div>
	</Modal>
</div>
