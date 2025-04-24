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
