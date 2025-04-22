<script lang="ts">
	import CpuStats from './CpuStats.svelte';
	import RamStats from './RamStats.svelte';
	import SpeedAnimation from './SpeedAnimation.svelte';

	let { systemInfo, showSpeedTest = $bindable() } = $props();

	function mbToGb(mb: number) {
		return (mb / 1024).toFixed(2);
	}

	function getInternetSpeed() {
		showSpeedTest = !showSpeedTest;
	}
</script>

<div
	class="bg-slate-200 dark:bg-slate-900 dark:text-white p-6 rounded-2xl flex flex-col w-[95vw] h-[80vh]"
>
	{#if showSpeedTest}
		<div>
			<button class="float-right -mt-2 cursor-pointer" onclick={getInternetSpeed}>Back</button>
			<SpeedAnimation />
		</div>
	{:else}
		<div class="grid grid-cols-2 items-center">
			<div>
				<div class="flex flex-col justify-center items-center space-y-2">
					{#if systemInfo.host_os == 'windows'}
						<img src="/windows.png" alt="windows" class="w-40 h-40" />
						<div class="text-xs">
							{systemInfo.host_platform} - {systemInfo.host_platform_version}
						</div>
					{/if}
				</div>
			</div>
			<div class="flex flex-col gap-2 items-start">
				<div class="grid grid-cols-2 justify-end items-center space-x-6 p-1">
					<CpuStats />
					<RamStats />
				</div>
				<div class="text-sm">
					<p>CPU Model: <span class="font-bold">{systemInfo.cpu_model}</span></p>
					<p>Vendor: <span class="font-bold">{systemInfo.cpu_vendor}</span></p>
					<p>Total Cores: <span class="font-bold">{systemInfo.cpu_cores} Cores</span></p>
					<p>
						Installed Memory: <span class="font-bold">{mbToGb(systemInfo.total_ram_mb)} GB</span>
					</p>
					<p>Disk Type: <span class="font-bold">{systemInfo.disk_type}</span></p>
					<p>Total Disk Space: <span class="font-bold">{systemInfo.total_disk_gb} GB</span></p>
					<p>Available Disk Space: <span class="font-bold">{systemInfo.free_disk_gb} GB</span></p>
					<p>Network Status: <span class="font-bold">{systemInfo.network_status}</span></p>
				</div>

				<div class="text-xs self-end mt-6">
					<button class="dark:bg-zinc-700 bg-zinc-300 p-2 cursor-pointer rounded-md"
						>Copy Stat</button
					>
					<button
						class="dark:bg-zinc-700 bg-zinc-300 p-2 cursor-pointer rounded-md"
						onclick={getInternetSpeed}>Check Internet Speed</button
					>
				</div>
			</div>
		</div>
	{/if}
</div>
