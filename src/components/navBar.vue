<template>
    <header class="mdl-layout__header">
        <div class="mdl-layout__header-row">      
			<router-link :to="{ name: 'main'}">
				<span class="mdl-layout-title">Parmyay</span>
			</router-link>
			<router-link :to="{ name: 'test'}" class="testPageLink">Test Page</router-link>
            <div class="mdl-layout-spacer"></div>
			<tri-switch class="switcher" :option1="option1" :option2="option2" :option3="option3" :option="option" v-on:changeOption="changeOption"/>
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
            TriSwitch
        }
    })
    export default class NavBar extends Vue {
        construct() {
        }

        option1 = true;
		option2 = false;
		option3 = false;
		option = "parmy";

		changeOption (): void {
            if(this.option1) {
				this.option1 = false;
				this.option2 = true;
				this.option3 = false;
				this.option = "parmi";
			} else if(this.option2) {
				this.option1 = false;
				this.option2 = false;
				this.option3 = true;
				this.option = "parma";
			} else {
				this.option1 = true;
				this.option2 = false;
				this.option3 = false;
				this.option = "parmy";
			}
			this.$emit("changeLang", this.option);
        };
    }
</script>