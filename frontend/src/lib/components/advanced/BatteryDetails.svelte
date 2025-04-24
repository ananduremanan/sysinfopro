<script lang="ts">
	import { GetBatteryDetails } from '$lib/wailsjs/go/main/App';
	import Card from '../ui/Card.svelte';
	import RefreshButton from '../ui/RefreshButton.svelte';

	let batteryDetails: any = $state(undefined);
	let isLoadingBattery = $state(false);

	async function getBatteryDetials() {
		try {
			isLoadingBattery = true;
			const battery = await GetBatteryDetails();
			if (battery) {
				batteryDetails = JSON.parse(battery);
			} else {
				batteryDetails = null;
			}
		} catch (error) {
			console.error('Failed to get battery details:', error);
		} finally {
			setTimeout(() => {
				isLoadingBattery = false;
			}, 500);
		}
	}

	$effect(() => {
		getBatteryDetials();
	});
</script>

<Card isLoading={isLoadingBattery}>
	<div class="h-28 flex justify-center items-center">
		<div class="absolute top-4 right-4">
			<RefreshButton onclick={getBatteryDetials} showText={false} />
		</div>
		{#if isLoadingBattery}
			<div class="">Loading Battery Details...</div>
		{:else}
			<div class="flex justify-center items-center h-full">
				<img src="/battery.png" alt="battery" class="w-24 h-24" />
				<div class="flex-1">
					<div class="font-bold mb-2">Battery Details</div>
					{#if batteryDetails}
						<div>
							<p>
								Percentage : <span class="font-bold">{batteryDetails.percentage.toFixed(0)}%</span>
							</p>
							<p>State : <span class="font-bold">{batteryDetails.state}</span></p>
							<p>
								Designed Capacity : <span class="font-bold"
									>{batteryDetails.design_capacity} mWh</span
								>
							</p>
							<p>
								Designed Capacity : <span class="font-bold"
									>{batteryDetails.current_capacity} mWh</span
								>
							</p>
						</div>
					{:else}
						<div class="font-semibold">No Battery Details Found</div>
						<p class="text-gray-500">
							Battery Details are only available to laptops and devices with a battery.
						</p>
					{/if}
				</div>
			</div>
		{/if}
	</div>
</Card>
