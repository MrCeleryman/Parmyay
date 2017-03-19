<template>
    <header class="mdl-layout__header">
        <div class="mdl-layout__header-row">      
			<router-link :to="{ name: 'main'}">
				<span class="mdl-layout-title">Parmyay</span>
			</router-link>
			<router-link :to="{ name: 'test'}" class="testPageLink">Test Page</router-link>
            <div class="mdl-layout-spacer"></div>
			<tri-switch></tri-switch>
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

	.testPageLink {
		padding-left: 50px;
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