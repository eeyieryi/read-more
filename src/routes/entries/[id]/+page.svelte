<script lang="ts">
	import EntryDetail from './EntryDetail.svelte';
	import type { PageServerData } from './$types';
	import { goto, invalidateAll } from '$app/navigation';
	import AudioPlayer from '$lib/components/AudioPlayer.svelte';
	import type { entryRepo } from '$lib/server/db';
	import { PUBLIC_STATIC_FILES_BASE_URL } from '$env/static/public';

	export let data: PageServerData;

	let showActions = false;

	const audioFilename =
		data.entryData.audioFilename.length > 0
			? `${PUBLIC_STATIC_FILES_BASE_URL}/audios/${data.entryData.audioFilename}`
			: '';

	const deleteEntry = async () => {
		if (
			!confirm(
				`Do you really want to delete ${data.entryData.title} entry?`
			)
		)
			return;
		const res = await fetch(`/api/entries/${data.entryData.id}`, {
			method: 'DELETE'
		});
		if (res.ok) {
			goto(`/collections/${data.entryData.collectionId}`);
		}
	};

	const cloneEntry = (): entryRepo.UpdateOneDto => {
		return {
			...data.entryData
		};
	};
	type ViewerMode = 'view' | 'edit';
	let currentMode: ViewerMode = 'view';
	let updateOneDto: entryRepo.UpdateOneDto;
	updateOneDto = cloneEntry();
	const saveEntry = async () => {
		const res = await fetch(`/api/entries/${updateOneDto.id}`, {
			method: 'PUT',
			body: JSON.stringify(updateOneDto),
			headers: {
				'content-type': 'application/json'
			}
		});
		if (res.ok) {
			await invalidateAll();
			currentMode = 'view';
			showActions = false;
		}
	};
	const discardChanges = async () => {
		updateOneDto = cloneEntry();
		currentMode = 'view';
		showActions = false;
	};

	const updateEntryStatus = async () => {
		const res = await fetch(`/api/entries/${data.entryData.id}`, {
			method: 'PUT',
			body: JSON.stringify(updateOneDto),
			headers: {
				'content-type': 'application/json'
			}
		});
		if (res.ok) {
			await invalidateAll();
			goto(`/collections/${data.entryData.collectionId}`);
		}
	};
</script>

<svelte:head>
	<title>{data.entryData.title} - {data.entryData.collection.title}</title>
</svelte:head>

{#if currentMode === 'view'}
	<div class="flex h-full flex-col overflow-auto text-center">
		<EntryDetail data="{data}" />
		<div
			class="align-center mb-2 ml-4 mt-auto flex items-center justify-center">
			{#if showActions}
				<div class="flex flex-row gap-x-4">
					<button
						class="block border px-3 py-2 text-lg"
						on:click|preventDefault="{deleteEntry}"
						>Delete Entry</button>
					<button
						class="ml-auto block border px-3 py-2 text-lg"
						on:click|preventDefault="{() => (currentMode = 'edit')}"
						>Edit Entry</button>
				</div>
				<button
					class="ml-auto px-3 py-2"
					on:click|preventDefault="{() =>
						(showActions = !showActions)}">-</button>
			{:else}
				<!-- TODO: Move AudioPlayer to somewhere else -->
				{#if audioFilename.length > 0}
					<AudioPlayer src="{audioFilename}" />
				{/if}
				<label>
					Mark as {data.entryData.seen ? 'unread' : 'read'}
					<input
						type="checkbox"
						bind:checked="{updateOneDto.seen}"
						on:change|preventDefault="{updateEntryStatus}" />
				</label>
				<button
					class="invisible ml-auto block border px-3 py-2 text-lg uppercase"
					>0</button>
				<button
					class="ml-auto px-3 py-2"
					on:click|preventDefault="{() =>
						(showActions = !showActions)}">+</button>
			{/if}
		</div>
	</div>
{:else if currentMode === 'edit'}
	<div class="m-4">
		<form
			class="flex h-full flex-col"
			on:submit|preventDefault="{saveEntry}">
			<div class="mb-4 flex flex-col">
				<label
					class="mb-2 text-lg font-bold uppercase text-gray-900"
					for="title">title</label>
				<input
					class="border px-3 py-2 text-gray-900"
					id="title"
					type="text"
					bind:value="{updateOneDto.title}" />
			</div>
			<div class="mb-4 flex flex-row items-center">
				<label
					class="mr-2 text-lg font-bold uppercase text-gray-900"
					for="collection">collection</label>
				<select
					class="px-3 py-2 text-gray-900"
					id="collection"
					bind:value="{updateOneDto.collectionId}">
					{#each data.collections as collection (collection.id)}
						<option
							value="{collection.id}"
							selected="{updateOneDto.collectionId ===
								collection.id}">{collection.title}</option>
					{/each}
				</select>
			</div>
			<div class="mb-4 flex flex-col">
				<label
					class="mb-2 text-lg font-bold uppercase text-gray-900"
					for="url">url</label>
				<input
					class="border px-3 py-2 text-gray-900"
					id="url"
					type="text"
					bind:value="{updateOneDto.url}" />
			</div>
			<div class="mb-8 flex flex-col">
				<label
					class="mb-2 text-lg font-bold uppercase text-gray-900"
					for="transcription">transcription</label>
				<textarea
					class="h-96 resize-none border px-3 py-2 text-gray-900"
					id="transcription"
					bind:value="{updateOneDto.transcription}"></textarea>
			</div>
			<div class="mt-auto flex flex-row justify-between">
				<button
					class="block border px-3 py-2 text-lg"
					on:click|preventDefault="{discardChanges}"
					type="button">Discard Changes</button>
				<button class="block border px-3 py-2 text-lg" type="submit"
					>Save Entry</button>
			</div>
		</form>
	</div>
{/if}
