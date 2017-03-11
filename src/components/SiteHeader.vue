<template>
    <header class="mdl-layout__header">
        <div class="mdl-layout__header-row">
            <!-- Title -->
            <span class="mdl-layout-title">Parmyay</span>
            <!-- Add spacer, to align navigation to the right -->
            <div class="mdl-layout-spacer"></div>
            <!-- Navigation. We hide it in small screens. -->
            <nav class="mdl-navigation mdl-layout--large-screen-only">
                Phrasing: 
                <label class="mdl-switch mdl-js-switch mdl-js-ripple-effect" for="language-switch">
                    <input type="checkbox" id="language-switch" class="mdl-switch__input" checked @change="changeLang">
                </label>
            </nav>
        </div>
    </header>
</template>
<style scoped>
    .mdl-layout--fixed-header {
        padding-bottom: 5px;
    }
</style>
<script lang="ts">
    import {Vue, Component} from 'av-ts'

    @Component
    export default class SiteHeader extends Vue {
        construct() {
        }

        changeLang = (e, a) => {
            this.$emit("changeLang", e.target.checked == true  ? "parmi" : "parma");
        }
    }
</script>