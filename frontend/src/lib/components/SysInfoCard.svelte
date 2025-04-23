<script lang="ts">
	import { checkInternetConnection, copyStat, mbToGb } from '$lib';
	import CpuStats from './CpuStats.svelte';
	import RamStats from './RamStats.svelte';
	import SpeedAnimation from './SpeedAnimation.svelte';

	let { systemInfo, showSpeedTest = $bindable() } = $props();

	let logMessage = $state<string | undefined>(undefined);

	async function getInternetSpeed() {
		const isConnected = await checkInternetConnection();

		if (!isConnected) {
			logMessage = 'No internet connection detected!';
			setTimeout(() => {
				logMessage = undefined;
			}, 2000);
			return;
		}

		showSpeedTest = !showSpeedTest;
	}

	async function copyStatHandler() {
		logMessage = await copyStat(systemInfo);

		setTimeout(() => {
			logMessage = undefined;
		}, 2000);
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
					{:else if systemInfo.host_os == 'darwin'}
						<img src="/mac.png" alt="macos" class="w-40 h-40" />
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
					<p>
						CPU Model: <span
							class="font-bold inline-block truncate max-w-[250px] align-bottom"
							title={systemInfo.cpu_model}>{systemInfo.cpu_model}</span
						>
					</p>
					<p>Vendor: <span class="font-bold">{systemInfo.cpu_vendor}</span></p>
					<p>Total Cores: <span class="font-bold">{systemInfo.cpu_cores} Cores</span></p>
					<p>
						Installed Memory: <span class="font-bold">{mbToGb(systemInfo.total_ram_mb)} GB</span>
					</p>
					<p>Disk Type: <span class="font-bold">{systemInfo.disk_type}</span></p>
					<p>Total Disk Space: <span class="font-bold">{systemInfo.total_disk_gb} GB</span></p>
					<p>Available Disk Space: <span class="font-bold">{systemInfo.free_disk_gb} GB</span></p>
					<p>Network Status: <span class="font-bold">{systemInfo.network_status}</span></p>
					<p>Logged In User: <span class="font-bold">{systemInfo.logged_in_user}</span></p>
				</div>

				<div class="text-xs self-end mt-2">
					<button
						class="dark:bg-zinc-700 bg-zinc-300 p-2 cursor-pointer rounded-md"
						onclick={copyStatHandler}>Copy Stat</button
					>
					<button
						class="dark:bg-zinc-700 bg-zinc-300 p-2 cursor-pointer rounded-md"
						onclick={getInternetSpeed}>Check Internet Speed</button
					>
					{#if logMessage}
						{#if logMessage === 'Stats copied to clipboard!'}
							<div class="text-xs mt-1 absolute text-green-500">{logMessage}</div>
						{:else}
							<div class="text-red-500 mt-1 absolute">{logMessage}</div>
						{/if}
					{/if}
				</div>
			</div>
		</div>
	{/if}
</div>
