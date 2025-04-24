<script lang="ts">
	import Card from '../ui/Card.svelte';

	let isLoadingBattery = $state(false);
</script>

<Card isLoading={isLoadingBattery}>
	<div class="h-full flex justify-center items-center">
		<div class="flex justify-center gap-2 items-center h-full">
			<div class="flex-1">
				<div class="font-bold mb-2">System Health Score</div>

				<p>Battery Details are only available to laptops and devices with a battery.</p>
			</div>
			<img src="/pc.png" alt="battery" class="w-24 h-24" />
		</div>
	</div>
</Card>
