<script lang="ts">
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
	import { PUBLIC_BACKEND_API } from '$env/static/public';
	import { entryModel } from '$lib/models';

	export let data: PageData;

	const createOneDto: entryModel.CreateOneDto = {
		title: '',
		description: '',
		url: '',
		transcription: '',
		audioFilename: '',
		collectionId: data.fromCollectionId ?? '',
		collectionTitle: ''
	};

	let filesToUpload: FileList | undefined;

	const handleSubmit = async (e: SubmitEvent) => {
		if (!e.defaultPrevented) {
			e.preventDefault();
		}

		if (filesToUpload !== undefined && filesToUpload.length === 1) {
			// if there is an audio file to upload
			const audioFile = filesToUpload[0];
			if (audioFile.name.length > 0 && createOneDto.title.length > 0) {
				const body = new FormData();
				body.set('audioFile', audioFile);
				body.set('entryTitle', createOneDto.title);
				const res = await fetch(`${PUBLIC_BACKEND_API}/upload`, {
					method: 'POST',
					body
				});
				if (res.ok) {
					const data = (await res.json()) as {
						audioFilename: string;
					};
					createOneDto.audioFilename = data.audioFilename;
				} else {
					// TODO: handle error
					return;
				}
			} else {
				// TODO: handle error
				return;
			}
		}

		const result = entryModel.validators.createOne(createOneDto);
		if (result.success) {
			const res = await fetch(`${PUBLIC_BACKEND_API}/entries`, {
				method: 'POST',
				body: JSON.stringify(createOneDto),
				headers: {
					'content-type': 'application/json'
				}
			});
			if (res.ok) {
				const createdEntry = (await res.json()) as entryModel.Entry;
				goto(`/entries/${createdEntry.id}`);
			} else {
				// TODO: handle error
				return;
			}
		} else {
			// TODO: handle error
			return;
		}
	};
</script>

<svelte:head>
	<title>New Entry</title>
</svelte:head>

<div class="flex h-full flex-col">
	<h1 class="mb-4 block text-xl font-bold uppercase text-gray-900">
		New Entry
	</h1>
	<form
		class="flex h-full flex-col"
		enctype="multipart/form-data"
		on:submit|preventDefault="{handleSubmit}">
		<div class="mb-4 flex flex-col">
			<label
				class="mb-2 text-lg font-bold uppercase text-gray-900"
				for="title">title</label>
			<input
				class="border px-3 py-2 text-gray-900"
				id="title"
				name="title"
				type="text"
				bind:value="{createOneDto.title}" />
		</div>
		<div class="mb-4 flex flex-col">
			<label
				class="mb-2 text-lg font-bold uppercase text-gray-900"
				for="url">url</label>
			<input
				class="border px-3 py-2"
				id="url"
				name="url"
				type="text"
				bind:value="{createOneDto.url}" />
		</div>
		<div class="mb-4 flex flex-col">
			<label
				class="mb-2 text-lg font-bold uppercase text-gray-900"
				for="url">audio file</label>
			<input
				class="border px-3 py-2"
				id="audioFile"
				name="audioFile"
				type="file"
				bind:files="{filesToUpload}"
				accept="audio/mp3 audio/ogg" />
		</div>
		<div class="mb-4 flex flex-col">
			<label
				class="mb-2 text-lg font-bold uppercase text-gray-900"
				for="transcription">transcription</label>
			<textarea
				class="h-56 resize-none border px-3 py-2 text-gray-900"
				id="transcription"
				name="transcription"
				bind:value="{createOneDto.transcription}"></textarea>
		</div>
		<div class="mb-8 flex w-full flex-row">
			<div class="mr-16 flex flex-row items-center">
				<label
					class="mr-8 text-lg font-bold uppercase text-gray-900"
					for="collection">existing collection</label>
				<select
					class="px-16 py-2 text-lg text-gray-900"
					id="collection"
					name="collection"
					bind:value="{createOneDto.collectionId}">
					<option value="">-</option>
					{#each data.collections as collection (collection.id)}
						<option value="{collection.id}"
							>{collection.title}</option>
					{/each}
				</select>
			</div>
			<div>
				<label
					class="mb-2 text-lg font-bold uppercase text-gray-900"
					for="new-collection">new collection</label>
				<input
					class="border px-3 py-2"
					id="new-collection"
					name="new-collection"
					type="text"
					bind:value="{createOneDto.collectionTitle}" />
			</div>
		</div>
		<div class="mt-auto flex flex-row">
			<button class="ml-auto block border px-3 py-2 text-lg" type="submit"
				>Create Entry</button>
		</div>
	</form>
</div>
