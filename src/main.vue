<template>
  <div class="mdl-layout mdl-js-layout mdl-layout--fixed-header">
    <nav-bar v-on:changeLang="changeLang"></nav-bar>
	<router-view v-bind:LANG="LANG"></router-view>
  </div>
</template>
<style scoped>
    #main {
        position: relative;
    }
</style>
<script lang="ts">
    import { Vue, Component } from "av-ts";
    import NavBar from "./components/navBar";

    @Component({
        components: {
            NavBar
        }
    })
    export default class Main extends Vue {
        LANG = {
            "parmy": "parmy"
        };

        changeLang = (newNoun) => {
            this.LANG.parmy = newNoun;
        };
    }
</script>
