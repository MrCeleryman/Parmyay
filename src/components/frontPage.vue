<template>
<main class="mdl-layout__content">
    <div class="page-content">
		<div class="main-text">
			<h1>Parmyay</h1>
        	<h2>Here to help you find your nearest {{ localisations.parma }}</h2>
		</div>
    </div>
  </main>
</template>
<style scoped>
    h1, h2 {
        color: green;
		text-align: center;
    }

	.main-text {
		position: absolute;
		top: 40%;
		left: 50%;
		transform: translate(-50%, -40%);
	}
</style>
<script lang="ts">
    import {Vue, Component, Watch} from "av-ts";
    import {LocalisedStrings} from "../util/localisedStrings";

    @Component
    export default class FrontPage extends Vue {
        localisations = LocalisedStrings.currentLocalisation;
    }
</script>