<template>
  <div class="mdl-layout mdl-js-layout mdl-layout--fixed-header">
    <site-header v-on:changeLang="changeLang"></site-header>
    <site-intro :LANG="LANG" ></site-intro>
  </div>
</template>
<style scoped>
    #main {
        position: relative;
    }
</style>
<script lang="ts">
    import { Vue, Component } from "av-ts";
    import SiteHeader from "./components/navBar";
    import SiteIntro from "./components/frontPage";

    @Component({
        components: {
            SiteHeader,
            SiteIntro
        }
    })
    export default class App extends Vue {
        LANG = {
            "parmy": "parmy"
        };

        changeLang = (newNoun) => {
            this.LANG.parmy = newNoun;
        };
    }
</script>
