<script lang="ts">
	let { showModal = $bindable(false), children, showClose = true } = $props();
</script>

{#if showModal}
	<div
		aria-labelledby="modal-title"
		role="dialog"
		aria-modal="true"
		class="fixed z-10 overflow-y-auto inset-0 flex items-center justify-center bg-slate-700/50 dark:bg-slate-700/90 transition-opacity"
	>
		<div class="bg-zinc-200 m-10 md:w-[80vh] rounded-xl dark:bg-zinc-800 dark:text-white">
			{#if showClose}
				<button class="float-right p-2 text-xs cursor-pointer" onclick={() => (showModal = false)}
					>Close</button
				>
			{/if}
			<div class="p-4 mt-2">{@render children()}</div>
		</div>
	</div>
{/if}
