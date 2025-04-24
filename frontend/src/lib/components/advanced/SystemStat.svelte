<script lang="ts">
	import Card from '../ui/Card.svelte';

	let { overallScore } = $props();
	let isLoadingBattery = $state(false);
</script>

<Card isLoading={isLoadingBattery}>
	<div class="h-full flex justify-center items-center">
		<div class="flex justify-center gap-2 items-center h-full">
			<div class="flex-1">
				<div class="font-bold mb-2">System Health Score</div>
				<div class="text-2xl font-bold text-green-500">{overallScore}</div>
				<p class="text-xs text-gray-400">Based on Battery Capacity, Available Disk Space and Installed Memory.</p>
			</div>
			<img src="/pc.png" alt="battery" class="w-24 h-24" />
		</div>
	</div>
</Card>
