<script lang="ts">
	import Credits from '$lib/components/ui/Credits.svelte';
	import Logo from '$lib/components/ui/Logo.svelte';
	import Modal from '$lib/components/ui/Modal.svelte';
	import RefreshButton from '$lib/components/ui/RefreshButton.svelte';
	import SysInfoCard from '$lib/components/basic/SysInfoCard.svelte';
	import { GetSysInfo } from '$lib/wailsjs/go/main/App';
	import { Quit } from '$lib/wailsjs/runtime/runtime';

	let systemInfo = $state(undefined);
	let error = $state('');
	let showSpeedTest = $state(false);
	let showAdvancedDashboard = $state(false);
	let showQuitConfirm = $state(false);
	let isLoading = $state(false);

	function toggleBack() {
		if (showSpeedTest) {
			showSpeedTest = false;
		} else if (showAdvancedDashboard) {
			showAdvancedDashboard = false;
		}
	}

	function toggleQuitModal() {
		showQuitConfirm = !showQuitConfirm;
	}

	async function getSysInfo() {
		isLoading = true;
		try {
			const result = await GetSysInfo();
			if (result) {
				systemInfo = JSON.parse(result);
			}
		} catch (err) {
			error = `Failed to get system information: ${err}`;
		} finally {
			setTimeout(() => {
				isLoading = false;
			}, 500);
		}
	}

	$effect(() => {
		getSysInfo();
	});
</script>

{#if isLoading}
	<div
		class="min-h-screen flex flex-col items-center justify-center dark:bg-slate-800 bg-slate-300"
	>
		<div class="flex flex-col items-center justify-center dark:text-white">
			<div class="flex"><Logo />™</div>
			<div class="text-xs">Loading...</div>
		</div>
	</div>
{:else}
	<section
		class="min-h-screen flex flex-col items-center justify-center dark:bg-slate-800 bg-slate-300"
	>
		<div class="flex flex-col items-start space-y-1">
			<div class="flex justify-between items-baseline w-full">
				<div class="flex dark:text-white">
					<Logo />™
				</div>
				<div class="flex space-x-6 items-center">
					{#if !showSpeedTest && !showAdvancedDashboard}
						<RefreshButton onclick={getSysInfo} />
					{/if}

					{#if showAdvancedDashboard || showSpeedTest}
						<button class="dark:text-slate-400 cursor-pointer text-sm mr-4" onclick={toggleBack}
							>Back</button
						>
					{/if}
					<button class="dark:text-slate-400 cursor-pointer text-sm mr-2" onclick={toggleQuitModal}
						>Close</button
					>
				</div>
			</div>

			{#if systemInfo}
				<SysInfoCard {systemInfo} bind:showSpeedTest bind:showAdvancedDashboard />
			{/if}

			<Credits />
		</div>

		<Modal showModal={showQuitConfirm} showClose={false}>
			<div class="flex flex-col justify-center items-center text-sm gap-2">
				<div>Are You Sure to Quit?</div>
				<div class="flex gap-2 mt-2">
					<button
						class="px-2 dark:bg-slate-600 rounded border border-red-500 cursor-pointer"
						onclick={Quit}>Yes</button
					>
					<button class="border px-2 rounded cursor-pointer" onclick={toggleQuitModal}>No</button>
				</div>
			</div>
		</Modal>
	</section>
{/if}
