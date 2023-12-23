import type { Successful, Unsuccessful } from './utils';
import {
	PUBLIC_BACKEND_API,
	PUBLIC_BACKEND_BASE_URL
} from '$env/static/public';

function getAudioFilenameUrl(audioFilename: string) {
	return `${PUBLIC_BACKEND_BASE_URL}/audios/${audioFilename}`;
}

type fetch = typeof fetch;

async function uploadFile(
	fetch: fetch,
	body: FormData
): Promise<(Successful & { audioFilename: string }) | Unsuccessful> {
	try {
		const res = await fetch(`${PUBLIC_BACKEND_API}/upload`, {
			method: 'POST',
			body
		});
		if (res.ok) {
			const data = (await res.json()) as { audioFilename: string };
			return {
				success: true,
				audioFilename: data.audioFilename
			};
		}
		return { success: false };
	} catch (err) {
		return {
			success: false
		};
	}
}

const uploadRepo = {
	getAudioFilenameUrl,
	uploadFile
};
export default uploadRepo;
