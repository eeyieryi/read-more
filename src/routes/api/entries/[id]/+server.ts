import { json } from '@sveltejs/kit';
import type { RequestHandler } from '@sveltejs/kit';
import { inputError, notFoundError } from '$lib/errors';
import { entryRepo } from '$lib/server/db';

export const GET: RequestHandler = async ({ params }) => {
	const result = entryRepo.validators.id(params.id);
	if (!result.success) {
		throw inputError();
	}

	const entry = await entryRepo.findOneById(result.data);
	if (!entry) {
		throw notFoundError();
	}

	return json(entry);
};

export const DELETE: RequestHandler = async ({ params }) => {
	const result = entryRepo.validators.id(params.id);
	if (!result.success) {
		throw inputError();
	}

	const deletedEntry = await entryRepo.deleteOneById(result.data);

	return json(deletedEntry);
};

export const PUT: RequestHandler = async ({ params, request }) => {
	const data = await request.json();

	const result = entryRepo.validators.updateOne({
		id: params.id,
		title: data.title,
		description: data.description,
		url: data.url,
		transcription: data.transcription,
		audioFilename: data.audioFilename,
		seen: data.seen,
		collectionId: data.collectionId
	});
	if (!result.success) {
		throw inputError();
	}

	const updatedEntry = await entryRepo.updateOne(result.data);

	return json(updatedEntry);
};
