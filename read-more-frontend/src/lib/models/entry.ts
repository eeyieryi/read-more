import { z } from 'zod';
import type { Collection } from './collection';
import type { Unknown } from '$lib/models/utils';

export type Entry = {
	id: string;
	title: string;
	description: string;
	url: string;
	transcription: string;
	audioFilename: string;
	seen: boolean;
	collectionId: string;
};

export type EntryData = Entry & {
	collection: Collection;
};

const idSchema = z.string().min(1);

const createOneSchemaRaw = z.object({
	title: z.string().min(1),
	description: z.optional(z.string()),
	url: z.optional(z.string()),
	transcription: z.string().min(1),
	audioFilename: z.optional(z.string()),
	collectionId: z.string(),
	collectionTitle: z.string()
});

const createOneSchema = createOneSchemaRaw.refine(
	({ collectionId, collectionTitle }) => {
		return collectionId.length > 0 || collectionTitle.length > 0;
	},
	{
		message:
			'One of the fields (collectionId or collectionTitle) must be provided'
	}
);

const updateOneSchema = createOneSchemaRaw
	.extend({
		id: idSchema,
		seen: z.optional(z.boolean())
	})
	.omit({ collectionTitle: true });

export type CreateOneDto = z.infer<typeof createOneSchema>;
export type UpdateOneDto = z.infer<typeof updateOneSchema>;

export const validators = {
	id: (id: unknown) => idSchema.safeParse(id),
	createOne: (data: Unknown<CreateOneDto>) => createOneSchema.safeParse(data),
	updateOne: (data: Unknown<UpdateOneDto>) => updateOneSchema.safeParse(data)
};
