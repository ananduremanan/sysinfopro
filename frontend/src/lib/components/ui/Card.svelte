<script lang="ts">
	let { children, isLoading = false } = $props();
</script>

<div
	class="relative text-xs {isLoading &&
		'animate-pulse'} justify-center bg-linear-to-r from-slate-200 to-slate-300 dark:from-slate-700 dark:to-slate-900 border border-slate-300 dark:border-slate-700 rounded-2xl w-full p-4 h-full"
>
	{@render children()}
</div>
