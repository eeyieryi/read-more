<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';

	const dispatch = createEventDispatcher();

	export let data: PageData;
</script>

<div class="flex flex-col">
	<h1 class="block py-4 text-4xl text-gray-900">
		{data.collectionData.title}
	</h1>
	<div class="flex flex-row">
		<button
			class="ml-auto mr-4 block bg-gray-700 px-2.5 py-1 text-sm text-white"
			on:click|preventDefault="{() =>
				goto(`/entries/new?fromCollection=${data.collectionData.id}`)}"
			>Add Entry</button>
		<button
			class="block bg-red-500 px-2.5 py-1 text-sm text-white"
			on:click|preventDefault="{() => dispatch('delete')}"
			>Delete Collection</button>
	</div>
</div>
