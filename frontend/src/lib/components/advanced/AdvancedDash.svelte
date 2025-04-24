<script lang="ts">
	import BatteryDetails from './BatteryDetails.svelte';
	import DiskCleaner from './DiskCleaner.svelte';
	import SystemStat from './SystemStat.svelte';
</script>

<div class="mt-6 grid grid-cols-2 gap-2 h-[65vh]">
	<div class="flex flex-col space-y-2">
		<BatteryDetails />
		<SystemStat />
	</div>
	<DiskCleaner />
</div>
