<script lang="ts">
	import { CheckCleanerPermissions, ScanSafeCleanableFiles } from '$lib/wailsjs/go/main/App';
	import Card from '../ui/Card.svelte';
	import PieDisk from './PieDisk.svelte';

	let isScanningFiles = $state(false);
	let files: any = $state(undefined);

	async function checkPermission() {
		try {
			const permission = await CheckCleanerPermissions();
			if (permission) {
				console.log('Permission granted:', permission);
			} else {
				console.log('Permission denied:', permission);
			}
		} catch (error) {
			console.error(error);
		}
	}

	async function scanFiles() {
		isScanningFiles = true;
		try {
			const fileResult: any = await ScanSafeCleanableFiles();
			if (fileResult) {
				files = fileResult;
				console.log('Scan result:', fileResult);
			} else {
				console.log('No files found to clean.');
			}
		} catch (error) {
			console.error('Error scanning files:', error);
		} finally {
			isScanningFiles = false;
		}
	}

	async function analyseStorageSpace() {
		await checkPermission();
		await scanFiles();
	}
</script>

<Card isLoading={false}>
	<div class="flex flex-col h-[60vh]">
		<div class="flex space-x-2 items-start">
			<div class="flex flex-col space-y-2 flex-1">
				<div class="font-bold">Disk Cleaner Utility</div>
				<div class="flex w-full">
					<div class="flex-1">
						<p class="">
							Clear Temporary Files and Browser Cache for free up some space from your disk.
							Clearing browser cache may be results in clearing browser history and cache.
						</p>
					</div>
				</div>
				{#if !files}
					<button
						class="mt-2 p-1 text-xs bg-blue-500 text-white rounded-lg cursor-pointer w-36 {isScanningFiles &&
							'animate-pulse'}"
						onclick={analyseStorageSpace}
						>{isScanningFiles ? 'Scanning...' : 'Analyse Storage Space'}</button
					>
				{/if}
			</div>
			<img src="/server-cleaning.png" alt="disk-cleaner" class="w-24 h-24" />
		</div>

		{#if isScanningFiles}
			<div
				class="flex items-center justify-center mt-6 animate-pulse dark:bg-slate-700 bg-slate-300 border border-slate-200 dark:border-slate-700 w-full h-20 rounded-2xl"
			>
				Scanning For Files...
			</div>
		{:else if files}
			<div
				class="mt-4 dark:bg-slate-700 bg-slate-200 border border-slate-300 dark:border-slate-600 w-full rounded-2xl p-3"
			>
				<PieDisk {files} />
			</div>
		{/if}
	</div>
</Card>
