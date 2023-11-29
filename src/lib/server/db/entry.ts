import { z } from 'zod';
import prisma, { type Entry, type EntryData } from './prisma';
import type { Unknown } from '$lib/utils';

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
		message: 'one of the fields must be provided'
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

export async function findOneById(id: string): Promise<EntryData | null> {
	return await prisma.entry.findUnique({
		where: { id },
		include: { collection: true }
	});
}

export async function createOne(createOneDto: CreateOneDto): Promise<Entry> {
	return await prisma.entry.create({
		data: {
			seen: false,
			title: createOneDto.title,
			description: createOneDto.description ?? '',
			url: createOneDto.url ?? '',
			transcription: createOneDto.transcription,
			audioFilename: createOneDto.audioFilename ?? '',
			collection: {
				connectOrCreate: {
					where: {
						id: createOneDto.collectionId
					},
					create: {
						title: createOneDto.collectionTitle,
						description: '', // TODO: ??
						url: '' // TODO: ??
					}
				}
			}
		},
		include: {
			collection: true
		}
	});
}

export async function updateOne(updateOneDto: UpdateOneDto): Promise<Entry> {
	return await prisma.entry.update({
		where: {
			id: updateOneDto.id
		},
		data: {
			title: updateOneDto.title,
			description: updateOneDto.description,
			url: updateOneDto.url,
			transcription: updateOneDto.transcription,
			audioFilename: updateOneDto.audioFilename,
			seen: updateOneDto.seen,
			collectionId: updateOneDto.collectionId
		}
	});
}

export async function deleteOneById(id: string): Promise<Entry> {
	return await prisma.entry.delete({
		where: { id }
	});
}
