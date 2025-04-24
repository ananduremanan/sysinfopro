import { CheckInternetConnection } from './wailsjs/go/main/App';

export async function checkInternetConnection() {
	const result = await CheckInternetConnection();
	if (result) {
		return true;
	} else {
		return false;
	}
}

export function mbToGb(mb: number) {
	return (mb / 1024).toFixed(2);
}

export function bytesToGb(kb: number) {
	return (kb / 1024 / 1024 / 1024).toFixed(2);
}

export async function copyStat(systemInfo: any): Promise<string> {
	const stats = JSON.stringify(systemInfo, null, 2);

	try {
		await navigator.clipboard.writeText(stats);
		return 'Stats copied to clipboard!';
	} catch (err) {
		return `Failed to copy: ${err}`;
	}
}

export function getSystemHealthScore(metrics: any): number {
	const {
		battery: { designedCapacity, currentCapacity },
		disk: { total: totalDisk, free: freeDisk },
		ram: { total: totalRam, available: availableRam }
	} = metrics;

	// Avoid division by zero
	if (designedCapacity <= 0 || totalDisk <= 0 || totalRam <= 0) return 0;

	const batteryHealth = (currentCapacity / designedCapacity) * 100;
	const diskHealth = (freeDisk / totalDisk) * 100;
	const ramHealth = (availableRam / totalRam) * 100;

	const score = (batteryHealth + diskHealth + ramHealth) / 3;

	return Math.round(score * 100) / 100;
}
