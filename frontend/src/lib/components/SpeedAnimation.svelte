<script lang="ts">
	import { checkInternetConnection } from '$lib';
	import { GetInterNetSpeed } from '$lib/wailsjs/go/main/App';
	import { onDestroy } from 'svelte';

	// State variables
	let speed = $state(0);
	let animationFrame: any = $state();
	let isRunning = $state(false);
	let progress = $state(0);
	let phase = $state('idle');
	let speedTestResult: any = $state(undefined);

	let logMessage = $state<string | undefined>(undefined);

	async function getInternetSpeed() {
		const isConnected = await checkInternetConnection();

		if (!isConnected) {
			logMessage = 'No internet connection detected!';
			setTimeout(() => {
				logMessage = undefined;
			}, 2000);
			return false;
		}

		return true;
	}

	// Configuration for the animation
	let animationConfig = $state({
		download: {
			duration: 8000,
			maxSpeed: 85,
			finalSpeed: 50
		},
		upload: {
			duration: 8000,
			maxSpeed: 45,
			finalSpeed: 32.7
		}
	});

	// Start the speed test animation
	async function startTest() {
		const isConnected = await getInternetSpeed();

		if (!isConnected) return;

		if (isRunning) return;

		isRunning = true;
		phase = 'download';
		progress = 0;
		speed = 0;

		// Start download test animation
		const downloadStartTime = Date.now();

		function animateDownload() {
			const elapsed = Date.now() - downloadStartTime;
			progress = Math.min(100, (elapsed / animationConfig.download.duration) * 100);

			// Calculate speed with some oscillation for realism
			const progressRatio = elapsed / animationConfig.download.duration;

			if (progressRatio < 0.8) {
				// Ramp up with oscillation
				speed =
					progressRatio *
					animationConfig.download.maxSpeed *
					(0.8 + 0.2 * Math.sin(elapsed / 300)) *
					(0.9 + 0.1 * Math.random());
			} else {
				// Stabilize towards final value
				speed = animationConfig.download.finalSpeed * (0.97 + 0.03 * Math.sin(elapsed / 200));
			}

			if (elapsed < animationConfig.download.duration || !speedTestResult) {
				animationFrame = requestAnimationFrame(animateDownload);
			} else {
				// Transition to upload test
				phase = 'upload';
				progress = 0;
				speed = 0;
				const uploadStartTime = Date.now();

				function animateUpload() {
					const upElapsed = Date.now() - uploadStartTime;
					progress = Math.min(100, (upElapsed / animationConfig.upload.duration) * 100);

					// Calculate upload speed with oscillation for realism
					const upProgressRatio = upElapsed / animationConfig.upload.duration;

					if (upProgressRatio < 0.7) {
						// Ramp up with oscillation
						speed =
							upProgressRatio *
							animationConfig.upload.maxSpeed *
							(0.7 + 0.3 * Math.sin(upElapsed / 400)) *
							(0.85 + 0.15 * Math.random());
					} else {
						// Stabilize towards final value
						speed = animationConfig.upload.finalSpeed * (0.95 + 0.05 * Math.sin(upElapsed / 250));
					}

					if (upElapsed < animationConfig.upload.duration || !speedTestResult) {
						animationFrame = requestAnimationFrame(animateUpload);
					} else {
						// Complete test
						phase = 'complete';
						isRunning = false;
					}
				}

				animationFrame = requestAnimationFrame(animateUpload);
			}
		}

		animationFrame = requestAnimationFrame(animateDownload);

		GetInterNetSpeed()
			.then((res) => {
				speedTestResult = JSON.parse(res);
				animationConfig.download.finalSpeed = speedTestResult.download_speed;
				animationConfig.upload.finalSpeed = speedTestResult.upload_speed;
			})
			.catch((err) => {
				console.error('Error fetching internet speed: ', err);
			});
	}

	function resetTest() {
		if (isRunning) {
			cancelAnimationFrame(animationFrame);
		}
		isRunning = false;
		phase = 'idle';
		speed = 0;
		progress = 0;
	}

	// Convert speed to needle rotation angle
	function getNeedleRotation(speed: any) {
		// Map speed to angle (-120 to 120 degrees for a 240-degree arc)
		const currentMaxSpeed = phase === 'upload' ? 50 : 100;
		const angle = -120 + (Math.min(speed, currentMaxSpeed) / currentMaxSpeed) * 240;
		return `rotate(${angle}deg)`;
	}

	// Get color based on speed
	function getSpeedColor(s: number) {
		if (phase === 'upload') {
			// Upload colors
			if (s < 5) return 'text-red-500';
			if (s < 15) return 'text-yellow-500';
			return 'text-green-500';
		} else {
			// Download colors
			if (s < 10) return 'text-red-500';
			if (s < 30) return 'text-yellow-500';
			return 'text-green-500';
		}
	}

	// Clean up on component destruction
	onDestroy(() => {
		if (animationFrame) {
			cancelAnimationFrame(animationFrame);
		}
	});
</script>

<!-- Speedometer -->
{#if phase === 'complete'}
	<div class="flex flex-col items-center justify-center gap-6 mt-20">
		<div class="flex gap-4">
			<div class="flex mb-2">
				<div class="font-medium text-6xl">
					{animationConfig.download.finalSpeed.toFixed(1)}<span class="text-sm">Mbps</span>
					<div class="text-sm text-gray-400">Download Speed</div>
				</div>
			</div>
			<div class="flex">
				<div class="font-medium text-6xl">
					{animationConfig.upload.finalSpeed.toFixed(1)}<span class="text-sm">Mbps</span>
					<div class="text-sm text-gray-400">Upload Speed</div>
				</div>
			</div>
		</div>

		<button
			onclick={startTest}
			class="px-4 text-sm py-2 bg-blue-500 text-white font-medium rounded-lg cursor-pointer"
		>
			{phase === 'complete' ? 'Test Again' : 'Start Test'}
		</button>
	</div>
{:else}
	<div class="grid grid-cols-2">
		<div class="flex flex-col items-center justify-center">
			<div class="relative w-64 h-64 mx-auto mb-4">
				<!-- Speed markings -->
				<div class="absolute inset-0">
					{#each Array(11) as _, i}
						<div
							class="absolute w-1 h-4 bg-gray-400"
							style="left: calc(50% - 0.5px); top: 8px; transform: rotate({-120 +
								i * 24}deg); transform-origin: center 120px;"
						></div>
					{/each}
				</div>

				<!-- Speed labels -->
				<div class="absolute text-xs text-gray-600" style="left: 15%; top: 70%;">0</div>
				<div class="absolute text-xs text-gray-600" style="left: 25%; top: 30%;">
					{phase === 'upload' ? '15' : '25'}
				</div>
				<div
					class="absolute text-xs text-gray-600"
					style="left: 50%; top: 15%; transform: translateX(-50%);"
				>
					{phase === 'upload' ? '20' : '50'}
				</div>
				<div class="absolute text-xs text-gray-600" style="right: 25%; top: 30%;">
					{phase === 'upload' ? '35' : '75'}
				</div>
				<div class="absolute text-xs text-gray-600" style="right: 15%; top: 70%;">
					{phase === 'upload' ? '50' : '100'}
				</div>

				<!-- Speed needle -->
				<div
					class="absolute left-1/2 top-2 w-1 h-32 bg-red-500 rounded-full -mt-2"
					style="transform: {getNeedleRotation(speed)}; transform-origin: center bottom;"
				></div>

				<!-- Center hub -->
				<div class="absolute left-1/2 top-1/2 w-6 h-6 rounded-full bg-gray-800 -ml-3 -mt-3"></div>

				<!-- Speed display -->
				<div class="absolute bottom-8 left-0 right-0 text-center">
					<div class={`text-3xl font-bold ${getSpeedColor(speed)}`}>
						{speed.toFixed(1)}
					</div>
					<div class="text-sm text-gray-500">
						{phase === 'upload' ? 'Mbps Upload' : 'Mbps Download'}
					</div>
				</div>
			</div>

			<!-- Status -->
			<div class="text-center text-gray-600 mb-4 text-sm">
				{#if phase === 'idle'}
					Ready to test your connection speed
				{:else if phase === 'download'}
					Testing download speed...
				{:else if phase === 'upload'}
					Testing upload speed...
				{:else}
					Test completed!
				{/if}
			</div>
		</div>

		<div class="flex flex-col justify-center items-end gap-4">
			<div class="text-xs text-justify">
				Test your internet speed in less than 30 seconds. The test typically uses between 40 to 100
				MB of data, though faster connections may use more. <br /> <br /> To perform the test, you'll
				be connected to Cloudflare. Your IP address will be shared with and processed by them according
				to their privacy policy. Cloudflare manages the test and makes all results publicly available
				to support internet research. These results include your IP address and speed data, but no personal
				information about you as a user is included.
			</div>

			<!-- Action button -->
			<div class="">
				{#if phase === 'idle' || phase === 'complete'}
					<div>
						<button
							onclick={startTest}
							class="px-4 text-sm py-2 bg-blue-500 text-white font-medium rounded-lg cursor-pointer"
						>
							{phase === 'complete' ? 'Test Again' : 'Start Test'}
						</button>
					</div>
				{:else}
					<button
						onclick={resetTest}
						class="px-4 py-2 bg-gray-500 text-white font-medium rounded-lg text-sm cursor-pointer"
					>
						Cancel
					</button>
				{/if}
			</div>
			<div class="text-xs text-red-500">{logMessage}</div>
		</div>
	</div>
{/if}
