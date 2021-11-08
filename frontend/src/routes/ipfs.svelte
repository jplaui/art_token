<script context="module">
	import { browser, dev } from '$app/env';

	// we don't need any JS on this page, though we'll load
	// it in dev so that we get hot module replacement...
	export const hydrate = dev;

	// ...but if the client-side router is already loaded
	// (i.e. we came here from elsewhere in the app), use it
	export const router = browser;

	// since there's no dynamic data here, we can prerender
	// it so that it gets served as a static asset in prod
	export const prerender = true;
</script>

<svelte:head>
	<title>ART Token</title>
</svelte:head>

<div class="content">
	<h1>Upload Collection into IPFS</h1>

	<p>
		Collection Management: Database entries, per user a table of collections
		Collection: Name, increment id, 
	</p>

	<pre>open</pre>

	<p>
		Select Images or Folder: add remove images to collection
	</p>

	<pre>open</pre>

	<p>
		Upload collection to IPFS address
	</p>

	<pre>open</pre>

</div>

<style>
	.content {
		width: 100%;
		max-width: var(--column-width);
		margin: var(--column-margin-top) auto 0 auto;
	}
</style>
