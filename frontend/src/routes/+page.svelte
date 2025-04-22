<script lang="ts">
	import Credits from '$lib/components/Credits.svelte';
	import RefreshButton from '$lib/components/RefreshButton.svelte';
	import SysInfoCard from '$lib/components/SysInfoCard.svelte';
	import { GetSysInfo } from '$lib/wailsjs/go/main/App';
	import { Quit } from '$lib/wailsjs/runtime/runtime';

	let systemInfo = $state(undefined);
	let error = $state('');
	let showSpeedTest = $state(false);

	async function getSysInfo() {
		try {
			const result = await GetSysInfo();
			if (result) {
				systemInfo = JSON.parse(result);
			}
		} catch (err) {
			error = `Failed to get system information: ${err}`;
		}
	}

	$effect(() => {
		getSysInfo();
	});
</script>

<section
	class="min-h-screen flex flex-col items-center justify-center dark:bg-slate-800 bg-slate-300"
>
	<div class="flex flex-col items-start space-y-1">
		<div class="flex justify-between items-baseline w-full">
			<div class="flex dark:text-white">
				<img src="/logo.svg" alt="logo" class="w-full h-10" />™
			</div>
			<div class="flex space-x-6 items-center">
				{#if !showSpeedTest}
					<RefreshButton onclick={getSysInfo} />
				{/if}
				<button class="dark:text-slate-400 cursor-pointer text-sm mr-2" onclick={Quit}>Close</button
				>
			</div>
		</div>

		{#if systemInfo}
			<SysInfoCard {systemInfo} bind:showSpeedTest />
		{/if}

		<Credits />
	</div>
</section>
