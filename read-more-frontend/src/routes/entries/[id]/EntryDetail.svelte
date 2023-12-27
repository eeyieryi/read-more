<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="scrollable flex flex-col items-center">
	<h1 class="py-4 text-2xl text-gray-900">{data.entryData.title}</h1>
	<a
		class="mb-4 text-lg font-light text-gray-900 underline underline-offset-4"
		href="/collections/{data.entryData.collection.id}"
		>{data.entryData.collection.title}</a>
	<pre
		class="mx-16 whitespace-pre-wrap pb-96 pt-14 text-2xl text-gray-700 lg:mx-72">{data
			.entryData.transcription}</pre>
</div>

<style>
	@font-face {
		font-family: 'NotoSansSC';
		src: url('/fonts/NotoSansSC-VariableFont_wght.ttf') format('ttf');
		font-weight: normal;
		font-style: normal;
	}
	pre {
		font-family: 'NotoSansSC', sans-serif;
	}
</style>
