<script lang="ts">
	import Modal from './Modal.svelte';
	import { BrowserOpenURL } from '$lib/wailsjs/runtime/runtime';

	let showModal = $state(false);
	let modalTitle: string | undefined = $state(undefined);
	let today = new Date();
	let year = today.getFullYear();

	function setModal(title = '') {
		modalTitle = title;
		showModal = !showModal;
	}
</script>

<div class="self-end text-sm dark:text-slate-600 flex space-x-4 mr-2">
	<button class="hover:dark:text-slate-400 cursor-pointer" onclick={() => setModal('About')}>
		About</button
	>
	<button class="hover:dark:text-slate-400 cursor-pointer" onclick={() => setModal('Credits')}
		>Credits</button
	>
	<button
		class="hover:dark:text-slate-400 cursor-pointer"
		onclick={() => {
			BrowserOpenURL('https://github.com/ananduremanan/sysinfopro');
		}}>Source</button
	>
</div>

<Modal bind:showModal>
	<div>
		<div>
			{modalTitle}
		</div>
		<div class="text-xs mt-2">
			{#if modalTitle === 'About'}
				<div>
					SysInfo Pro is a lightweight, cross-platform system diagnostics tool designed to give you
					instant insights into your machine's performance and hardware configuration. Whether
					you're a developer, IT professional, or just a curious user, SysInfo Pro provides
					real-time data in a clean, readable format — with zero bloat.
					<div class="flex flex-col items-start space-x-2 mt-2">
						<div>Bug Report URL:</div>
						<button
							class="text-blue-500 hover:text-blue-700 cursor-pointer"
							onclick={() => {
								BrowserOpenURL('https://github.com/ananduremanan/sysinfopro/issues');
							}}>https://github.com/ananduremanan/sysinfopro/issues</button
						>
					</div>

					<div class="mt-4">Sysinfo Pro © {year} - Present</div>
					<div class="mt-1">sysinfopro.netlify.app</div>
				</div>
			{:else if modalTitle === 'Credits'}
				<div class="flex justify-between items-center">
					<div>
						<p>SysInfo Pro version 0.1</p>
						<p>Author: Nada Labs</p>
						<p>Icon Credit: Flaticon</p>
						<p>Made with Wails, Go, Svelte and Tailwind</p>
						<p class="mt-4">www.sysinfopro.in</p>
					</div>
					<div class="flex flex-col space-y-2 items-center justify-center">
						<img src="/Open_Source_Initiative.svg" alt="oss" class="w-20 h-20" />
						<div class="text-center">We love and Support <br /> Open Source</div>
					</div>
				</div>
			{/if}
		</div>
	</div>
</Modal>
