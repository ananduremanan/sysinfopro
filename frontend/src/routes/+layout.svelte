<script lang="ts">
	import '../app.css';

	let { children } = $props();
</script>

<main class="flex flex-col min-h-screen bg-gray-100">
	{@render children()}
</main>
