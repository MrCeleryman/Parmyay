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
    import {Vue, Component} from 'av-ts'
    import SiteHeader from './components/SiteHeader'
    import SiteIntro from './components/SiteIntro'

    @Component({
        components: {
            SiteHeader,
            SiteIntro
        }
    })
    export default class Entry extends Vue {
        LANG = {
            "parmi": "parmi"
        };

        changeLang = (newNoun) => {
            this.LANG.parmi = newNoun;
        }
    }
</script>
