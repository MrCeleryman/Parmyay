<template>
    <div class="winner">
        <h1>This page is supposed to look bad</h1>
    </div>
</template>
<style scoped>
    .winner h1 {
        color: green;
    }
</style>
<script lang="ts">
    import {
        Vue,
        Component
    } from 'av-ts'
    @Component
    export default class WinnerIsYou extends Vue {
        construct() {}
    }
</script>