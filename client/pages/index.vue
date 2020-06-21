<template>
  <div>
    <v-app>
    <div class="content-wrapper">
      <v-bottom-navigation
        :value="current"
        grow
        color="teal"
      >
        <v-btn
          v-for="data in datas"
          :key="data.id"
          @click="changeCharacter(data)"
        >
          <span>{{ data.name }}</span>
          <!-- <v-icon>mdi-history</v-icon> -->
        </v-btn>
      </v-bottom-navigation>
      <v-expansion-panels
        accordion
        class="panel-wrapper"
      >
        <v-expansion-panel>
          <v-expansion-panel-header>ダイス</v-expansion-panel-header>
          <v-expansion-panel-content>
            <AbilityCard :data="current"/>
            <AbilityCard :data="current"/>
            <AbilityCard :data="current"/>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
      <v-expansion-panels
        accordion
        class="panel-wrapper"
      >
        <v-expansion-panel>
          <v-expansion-panel-header>基本能力</v-expansion-panel-header>
          <v-expansion-panel-content>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
      <v-expansion-panels
        accordion
        class="panel-wrapper"
      >
        <v-expansion-panel>
          <v-expansion-panel-header>Item</v-expansion-panel-header>
          <v-expansion-panel-content>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </div>
    <!-- /.content-wrapper -->
    <div class="side-wrapper">

    </div>
    <!-- /.side-wrapper -->
    </v-app>
  </div>
</template>

<script>
import AbilityCard from "~/components/ui/AbilityCard";

export default {
  components: {
    AbilityCard
  },
  data() {
    return {
      current: {
        name: "Taro"
      },
      datas: [
        {
          name: "Taro"
        },
        {
          name: "Taro"
        },
        {
          name: "Taro"
        },
        {
          name: "Sato"
        }
      ]
    }
  },
  methods: {
    changeCharacter: function(data) {
      this.current = data
      console.log(this.current)
    }
  }
}
</script>

<style lang="scss" scoped>
.content-wrapper {
  width: 80%;
  max-width: 1000px;
  min-height: 100vh;
  border-radius: 10px;
  margin: 0 auto;
}

.panel-wrapper {
  margin: 30px auto;
}
</style>