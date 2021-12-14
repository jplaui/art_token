<!-- <script>
	import { browser, dev } from '$app/env';

	// we don't need any JS on this page, though we'll load
	// it in dev so that we get hot module replacement...
	export const hydrate = dev;

	// ...but if the client-side router is already loaded
	// (i.e. we came here from elsewhere in the app), use it
	export const router = browser;

	// since there's no dynamic data here, we can prerender
	// it so that it gets served as a static asset in prod
	export const prerender = false;
</script> -->

<script>
	import { goto } from '$app/navigation';

	let collectionImages = [{"image_id": "/blocks.jpg", "name": "collection name", "end_date": "2", "days_exist": "20"}, {"image_id": "/vr.jpg", "name": "collection name", "end_date": "50", "days_exist": "3"}, {"image_id": "/world.jpg", "name": "collection name", "end_date": "30", "days_exist": "4"}, {"image_id": "/art1.jpg", "name": "collection name", "end_date": "44", "days_exist": "3"}, {"image_id": "/pixel.jpg", "name": "collection name", "end_date": "12", "days_exist": "10"}];

	function enterAuction(image_index) {
		let url = "/collections/"+image_index.toString();
		console.log("url:", url);
		goto(url).catch(function(err) {
			console.log(err);
		});
	}
</script>

<svelte:head>
	<title>ART Token</title>
</svelte:head>

	<!-- Page Content -->
	<div class="w3-padding-large" id="main">
	  <!-- Header/Home -->
	  <header class="w3-container w3-padding-32 w3-center w3-black">
	    <h1 class="w3-jumbo"><span class="w3-hide-small">Demo:</span> ART Token</h1>
	    <p>Collections.</p>
	  </header>

	  <!-- Collections Section -->
		<div class="w3-content w3-justify w3-black">
			<p><a href="/">Back.</a></p>
		</div>

	  <div class="w3-content w3-justify w3-text-grey w3-padding-64">
	    <h2 class="w3-text-light-grey">Active Auctions</h2>
	    <hr style="width:200px" class="w3-opacity">
	    <p>Select a collection, connect your wallet, and start minting your NFTs.
	    </p>
	    	<!-- Product grid -->
		  <div class="w3-row w3-grayscale">
		  	{#each collectionImages as { image_id, name, end_date, days_exist }, i}
		    <div class="w3-col l3 s6">
		      <div class="w3-container">
		      	<div class="w3-display-container">
		          <img alt="collectionItem" src="{image_id}" style="width:100%">
		          {#if days_exist <= 5}
		          	<span class="w3-tag w3-display-topleft">New</span>
		          {/if}
		          <div class="w3-display-middle w3-display-hover">
		            <button class="w3-button w3-black" on:click={() => enterAuction(i)}>Enter Auction <i class="fa fa-shopping-cart"></i></button>
		          </div>
		        </div>
		        <p>{name}<br><b>Ending in {end_date} days.</b></p>
		      </div>
		    </div>

		    

		    {/each}
		  </div>

	  <!-- End Collections Section -->
	  </div>
	</div>

<style>
	body, h1,h2,h3,h4,h5,h6 {font-family: "Montserrat", sans-serif}
	.w3-row-padding img {margin-bottom: 12px}
	/* Set the width of the sidebar to 120px */
	.w3-sidebar {width: 120px;background: #222;}
	/* Add a left margin to the "page content" that matches the width of the sidebar (120px) */
	#main {margin-left: 120px}
	/* Remove margins from "page content" on small screens */
	@media only screen and (max-width: 600px) {#main {margin-left: 0}}
</style>
