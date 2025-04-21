<script lang="ts">
	import SpeedAnimation from './SpeedAnimation.svelte';

	let { systemInfo } = $props();

	let showSpeedTest = $state(false);

	function mbToGb(mb: number) {
		return (mb / 1024).toFixed(2);
	}

	function getInternetSpeed() {
		showSpeedTest = !showSpeedTest;
	}
</script>

<div class="bg-slate-200 dark:bg-slate-900 dark:text-white p-6 rounded-2xl">
	{#if showSpeedTest}
		<div class="mt-6 w-[50vw]">
			<button class="float-right -mt-8 cursor-pointer" onclick={getInternetSpeed}>Back</button>
			<SpeedAnimation />
		</div>
	{:else}
		<div class="grid grid-cols-2 mt-6">
			<div class="flex flex-col justify-center items-center space-y-2">
				{#if systemInfo.host_os == 'windows'}
					<img src="/windows.png" alt="windows" class="w-40 h-40" />
					<div class="text-xs">{systemInfo.host_platform} - {systemInfo.host_platform_version}</div>
				{/if}
			</div>
			<div class="text-sm">
				<p>CPU Model: <span class="font-bold">{systemInfo.cpu_model}</span></p>
				<p>Vendor: <span class="font-bold">{systemInfo.cpu_vendor}</span></p>
				<p>Total Cores: <span class="font-bold">{systemInfo.cpu_cores} Cores</span></p>
				<p>
					Installed Memory: <span class="font-bold">{mbToGb(systemInfo.total_ram_mb)} GB</span>
				</p>
				<p>Free Memory: <span class="font-bold">{mbToGb(systemInfo.free_ram_mb)} GB</span></p>
				<p>Disk Type: <span class="font-bold">{systemInfo.disk_type}</span></p>
				<p>Total Disk Space: <span class="font-bold">{systemInfo.total_disk_gb} GB</span></p>
				<p>Available Disk Space: <span class="font-bold">{systemInfo.free_disk_gb} GB</span></p>
				<p>Network Status: <span class="font-bold">{systemInfo.network_status}</span></p>
			</div>
		</div>

		<div class="text-xs float-right mt-6">
			<button class="bg-zinc-700 p-2 cursor-pointer rounded-md">Copy Stat</button>
			<button class="bg-zinc-700 p-2 cursor-pointer rounded-md" onclick={getInternetSpeed}
				>Check Internet Speed</button
			>
		</div>
	{/if}
</div>
