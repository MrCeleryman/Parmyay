<template>
  <div class="mdl-layout mdl-js-layout mdl-layout--fixed-header">
    <nav-bar></nav-bar>
    <front-page></front-page>
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
    import FrontPage from "./components/frontPage";

    @Component({
        components: {
            NavBar,
            FrontPage
        }
    })
    export default class Main extends Vue {
    }
</script>
