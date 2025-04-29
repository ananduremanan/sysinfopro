<script lang="ts">
	import Modal from '$lib/components/ui/Modal.svelte';
	import { CheckUpdate } from '$lib/wailsjs/go/main/App';
	import '../app.css';

	let { children } = $props();

	const downloadUrl =
		'https://github.com/ananduremanan/sysinfopro/releases/download';

	let showUpdate = $state(false);
	let versionName = $state('');

	async function checkUpdate() {
		try {
			const result = await CheckUpdate();
			if (result) {
				console.log(result);
				showUpdate = JSON.parse(result).is_update_available;
			}
		} catch (error) {
			console.error('Error checking for updates:', error);
		}
	}

	$effect(() => {
		checkUpdate();
	});
</script>

<main class="flex flex-col min-h-screen bg-gray-100">
	{@render children()}

	<Modal showModal={showUpdate} showClose={false}>
		<div class="flex flex-col justify-center items-center text-xs gap-1">
			<div class="font-bold">Updates Available</div>
			<div>A New Version of SysInfoPro is Available</div>
			<div class="mt-2">
				<button class="px-2 py-1 bg-green-500 rounded-lg text-white cursor-pointer">
					<a href={`${downloadUrl}/${versionName}/sysinfopro.exe`}>Download Now</a></button
				>
				<button
					class="px-2 py-1 border border-green-500 rounded-lg cursor-pointer"
					onclick={() => {
						showUpdate = false;
					}}>Not Now</button
				>
			</div>
		</div>
	</Modal>
</main>
