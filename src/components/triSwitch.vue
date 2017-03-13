<template>
	<div class="doughnut" v-on:click="changeOption">
		<div class="switch" v-bind:class="{ option1: option1, option2: option2, option3: option3 }">
		</div>
		<p class="choice choice1">Parmy</p>
		<p class="choice choice2">Parmi</p>
		<p class="choice choice3">Parma</p>
	</div>
</template>
<style scoped>
    .doughnut {
		cursor: pointer;
		border: 7px solid rgba(66, 66, 66, 0.258824);
		border-radius: 100px;
		height: 25px;
		width: 25px;
	}

	.switch {
		position: relative;
		background-color: #fafafa;
		border-radius: 100px;
		height: 22px;
		width: 22px;
		transition: all .3s ease;
		box-shadow: 0 2px 2px 0 rgba(0,0,0,.14),
		0 3px 1px -2px rgba(0,0,0,.2),
		0 1px 5px 0 rgba(0,0,0,.12);
	}

	.option1 {
		left: 15px;
		top: 10px;
	}

	.option2 {
		left: -12px;
		top: 10px;
	}

	.option3 {
		left: 1px;
		top: -11px;
	}

	.choice {
		position: relative;
		font-size: 16px;
	}

	.choice1 {
		left: 40px;
		top: 10px;
	}

	.choice2 {
		left: -50px;
		top: -30px;
	}

	.choice3 {
		left: -10px;
		top: -140px;
	}

</style>
<script lang="ts">
    import {Vue, Component, Prop} from "av-ts";

    @Component({
		data: {
			option1: this.option1,
			option2: this.option2,
			option3: this.option3
		}
	})
    export default class TriSwitch extends Vue {
		@Prop
	    option1: boolean;
		@Prop
		option2: boolean;
		@Prop
		option3: boolean;
		@Prop
		option: string;

        construct() {};

		changeOption (): void {
			this.$emit("changeOption");
		}
    }
</script>