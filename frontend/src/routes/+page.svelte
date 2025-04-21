<script lang="ts">
	import Credits from '$lib/components/Credits.svelte';
	import SysInfoCard from '$lib/components/SysInfoCard.svelte';
	import { GetSysInfo } from '$lib/wailsjs/go/main/App';

	let systemInfo = $state(undefined);
	let error = $state('');

	$effect(() => {
		GetSysInfo()
			.then((res) => {
				systemInfo = JSON.parse(res);
			})
			.catch((err) => {
				error = err;
			});
	});
</script>

<section
	class="min-h-screen flex flex-col items-center justify-center dark:bg-slate-800 bg-slate-300"
>
	<div class="flex flex-col items-start space-y-1">
		<div class="flex dark:text-white"><img src="/logo.svg" alt="logo" class="w-full h-10" />™</div>

		{#if systemInfo}
			<SysInfoCard {systemInfo} />
		{/if}

		<Credits />
	</div>
</section>
