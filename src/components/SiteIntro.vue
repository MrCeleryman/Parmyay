<template>
<main class="mdl-layout__content">
    <div class="page-content">
        <h1>The current lang is: {{ LANG.parmi }}</h1>
    </div>
  </main>
</template>
<style scoped>
    h1 {
        color: green;
    }
</style>
<script lang="ts">
    import {Vue, Component, Prop} from "av-ts";
    
    @Component
    export default class SiteIntro extends Vue {
        @Prop
        LANG;

        construct() {};
    }
</script>