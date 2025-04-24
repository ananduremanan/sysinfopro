<script lang="ts">
	import { GetMemoryStats } from '$lib/wailsjs/go/main/App';

	type MemoryStats = {
		totalRAMMB: number;
		freeRAMMB: number;
		usedRAMMB: number;
		usagePercent: number;
		timestamp: number;
	};

	let memStats: MemoryStats | null = $state(null);
	let unsubscribe: (() => void) | null = $state(null);

	function formatRAM(mb: number): string {
		if (mb > 1024) {
			return (mb / 1024).toFixed(2) + ' GB';
		}
		return mb.toFixed(0) + ' MB';
	}

	function updateStats(data: MemoryStats) {
		memStats = data;
	}

	async function init() {
		try {
			const initialStats = await GetMemoryStats();
			updateStats(initialStats);

			// @ts-ignore
			unsubscribe = window.runtime.EventsOn('ram-stats-update', (data: MemoryStats) => {
				updateStats(data);
			});
		} catch (error) {
			console.error('Failed to get RAM stats:', error);
		}
	}

	$effect(() => {
		init();
	});

	$effect(() => {
		return () => {
			if (unsubscribe) {
				unsubscribe();
			}
		};
	});

	function getUsageColor(percent: number): string {
		if (percent < 50) return 'stroke-green-500';
		if (percent < 80) return 'stroke-yellow-500';
		return 'stroke-red-500';
	}

	const circumference = 2 * Math.PI * 45; // r = 45
</script>

<div class="">
	{#if memStats}
		<!-- Circular Arc -->
		<div class="flex items-center text-xs">
			<div class="relative w-14 h-14 mx-auto">
				<svg class="w-full h-full transform -rotate-90" viewBox="0 0 100 100">
					<circle cx="50" cy="50" r="45" class="stroke-zinc-300" stroke-width="10" fill="none" />
					<circle
						cx="50"
						cy="50"
						r="45"
						stroke-width="10"
						fill="none"
						stroke-linecap="round"
						stroke-dasharray={circumference}
						stroke-dashoffset={circumference - (circumference * memStats.usagePercent) / 100}
						class={`transition-all duration-700 ease-out ${getUsageColor(memStats.usagePercent)}`}
					/>
				</svg>
				<div class="absolute inset-0 flex items-center justify-center font-semibold text-xs">
					{memStats.usagePercent.toFixed(1)}%
				</div>
			</div>
			<div class="ml-1">
				<div>Total Ram: {formatRAM(memStats.totalRAMMB)}</div>
				<div>Used Ram: {formatRAM(memStats.usedRAMMB)}</div>
				<div>Free Ram: {formatRAM(memStats.freeRAMMB)}</div>
			</div>
		</div>
	{:else}
		<div class="text-gray-500 text-sm text-center">Loading RAM stats...</div>
	{/if}
</div>
