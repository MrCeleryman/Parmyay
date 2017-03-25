<template>
<div>
    <md-button>Default</md-button>
    <md-button class="md-primary">Primary</md-button>
    <md-button class="md-accent">Accent</md-button>
    <md-button class="md-warn">Warn</md-button>
    <md-button class="md-primary" disabled>Disabled</md-button>
    <md-button class="md-dense">Dense</md-button>
</div>
  <!--<div class="mdl-layout mdl-js-layout mdl-layout--fixed-header">
    <nav-bar></nav-bar>
	<router-view></router-view>
  </div>-->
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
