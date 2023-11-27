import { json } from '@sveltejs/kit';
import type { RequestHandler } from '@sveltejs/kit';
import { inputError } from '$lib/errors';
import { entryRepo } from '$lib/server/db';

export const POST: RequestHandler = async ({ request }) => {
	const data = await request.json();

	const result = entryRepo.validators.createOne({
		title: data.title,
		description: data.description,
		url: data.url,
		transcription: data.transcription,
		audioFilename: data.audioFilename,
		collectionId: data.collectionId,
		collectionTitle: data.collectionTitle
	});
	if (!result.success) {
		throw inputError();
	}

	const createdEntry = await entryRepo.createOne(result.data);

	return json(createdEntry);
};
