<script lang="ts">
	import { goto } from '$app/navigation';
	import { PUBLIC_BACKEND_API } from '$env/static/public';
	import type { collectionModel } from '$lib/models';

	const createOneDto: collectionModel.CreateOneDto = {
		title: '',
		description: '',
		url: ''
	};

	const handleSubmit = async () => {
		const res = await fetch(`${PUBLIC_BACKEND_API}/collections`, {
			method: 'POST',
			body: JSON.stringify(createOneDto),
			headers: {
				'content-type': 'application/json'
			}
		});
		if (res.ok) {
			const createdCollection =
				(await res.json()) as collectionModel.Collection;
			goto(`/collections/${createdCollection.id}`);
		}
	};
</script>

<svelte:head>
	<title>New Collection</title>
</svelte:head>

<div class="flex h-full flex-col">
	<h1 class="mb-4 block text-xl font-bold uppercase text-gray-900">
		New Collection
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
		<div class="mb-8 flex flex-col">
			<label
				class="mb-2 text-lg font-bold uppercase text-gray-900"
				for="url">url</label>
			<input
				class="border px-3 py-2 text-gray-900"
				id="url"
				type="text"
				bind:value="{createOneDto.url}" />
		</div>
		<div class="mt-auto flex flex-row">
			<button class="ml-auto border px-3 py-2 text-lg" type="submit"
				>Create Collection</button>
		</div>
	</form>
</div>
