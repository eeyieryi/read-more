<script lang="ts">
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
	import type { entryRepo, Entry } from '$lib/server/db';

	export let data: PageData;

	const createOneDto: entryRepo.CreateOneDto = {
		title: '',
		description: '',
		url: '',
		transcription: '',
		audioFilename: '',
		collectionId: data.fromCollectionId ?? '',
		collectionTitle: ''
	};

	const handleSubmit = async () => {
		const res = await fetch('/api/entries', {
			method: 'POST',
			body: JSON.stringify(createOneDto),
			headers: {
				'content-type': 'application/json'
			}
		});
		if (res.ok) {
			const createdEntry = (await res.json()) as Entry;
			goto(`/entries/${createdEntry.id}`);
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
		on:submit|preventDefault="{handleSubmit}">
		<div class="mb-4 flex flex-col">
			<label
				class="mb-2 text-lg font-bold uppercase text-gray-900"
				for="title">title</label>
			<input
				class="border px-3 py-2 text-gray-900"
				id="title"
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
				type="text"
				bind:value="{createOneDto.url}" />
		</div>
		<div class="mb-4 flex flex-col">
			<label
				class="mb-2 text-lg font-bold uppercase text-gray-900"
				for="transcription">transcription</label>
			<textarea
				class="h-56 resize-none border px-3 py-2 text-gray-900"
				id="transcription"
				bind:value="{createOneDto.transcription}"></textarea>
		</div>
		<div class="mb-4 flex flex-row items-center">
			<label
				class="mr-2 text-lg font-bold uppercase text-gray-900"
				for="collection">existing collection</label>
			<select
				class="px-3 py-2 text-lg text-gray-900"
				id="collection"
				bind:value="{createOneDto.collectionId}">
				<option value="">-</option>

				{#each data.collections as collection (collection.id)}
					<option value="{collection.id}">{collection.title}</option>
				{/each}
			</select>
		</div>
		<div class="mb-8 flex flex-col">
			<label
				class="mb-2 text-lg font-bold uppercase text-gray-900"
				for="new-collection">new collection</label>
			<input
				class="border px-3 py-2"
				id="new-collection"
				type="text"
				bind:value="{createOneDto.collectionTitle}" />
		</div>
		<div class="mt-auto flex flex-row">
			<button class="ml-auto block border px-3 py-2 text-lg" type="submit"
				>Create Entry</button>
		</div>
	</form>
</div>
