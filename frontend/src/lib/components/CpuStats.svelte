<script lang="ts">
	import { GetCPUStats } from '../wailsjs/go/main/App';

	// Define the structure of CPU stats data
	interface CPUStats {
		usage: number;
		timestamp: number;
	}

	let cpuUsage = $state(0);
	let timestamps = $state<string[]>([]);
	let usageHistory = $state<number[]>([]);
	const MAX_HISTORY_LENGTH = 30;
	let unsubscribe = $state<(() => void) | undefined>(undefined);

	// Format the CPU usage as a percentage with one decimal place
	let formattedCpuUsage = $derived(cpuUsage.toFixed(1));

	// Calculate color based on CPU usage
	let cpuColor = $derived(
		cpuUsage > 80 ? 'text-red-500' : cpuUsage > 60 ? 'text-amber-500' : 'text-green-500'
	);

	function updateStats(stats: CPUStats) {
		cpuUsage = stats.usage;

		// Add current data point to history arrays
		timestamps = [...timestamps, new Date(stats.timestamp * 1000).toLocaleTimeString()];
		usageHistory = [...usageHistory, stats.usage];

		// Keep history at a reasonable size
		if (timestamps.length > MAX_HISTORY_LENGTH) {
			timestamps = timestamps.slice(1);
			usageHistory = usageHistory.slice(1);
		}
	}

	async function init() {
		try {
			// Get initial CPU stats
			const initialStats = await GetCPUStats();
			updateStats(initialStats);

			// TypeScript ignore for window.runtime
			// @ts-ignore
			unsubscribe = window.runtime.EventsOn('cpu-stats-update', (data: CPUStats) => {
				updateStats(data);
			});
		} catch (error) {
			console.error('Failed to get CPU stats:', error);
		}
	}

	$effect(() => {
		init();
	});

	// Cleanup when component is destroyed
	$effect(() => {
		return () => {
			if (unsubscribe) {
				unsubscribe();
			}
		};
	});
</script>

<div class="flex flex-col gap-2">
	<div class="text-xs">CPU Utilization</div>
	<div class="{cpuColor} fill-current text-5xl">
		{formattedCpuUsage}%
	</div>
</div>
