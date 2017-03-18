<template>
    <header class="mdl-layout__header">
        <div class="mdl-layout__header-row">
            <span class="mdl-layout-title">Parmyay</span>
            <div class="mdl-layout-spacer"></div>
			<tri-switch class="switcher"></tri-switch>
        </div>
    </header>
</template>
<style scoped>
    .mdl-layout--fixed-header {
        padding-bottom: 5px;
    }

	.mdl-layout__header-row {
		height: 120px;
	}

	.switcher {
		margin-right: 50px;
	}

</style>
<script lang="ts">
    import { Vue, Component } from "av-ts";
    import TriSwitch from "./triSwitch";

    @Component({
        components: {
            "tri-switch": TriSwitch
        }
    })
    export default class NavBar extends Vue {

    }
</script>