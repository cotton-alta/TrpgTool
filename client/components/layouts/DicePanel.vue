<template>
  <div class="panel-wrapper">
        <!-- class="panel-wrapper" -->
      <v-expansion-panels
        accordion
      >
        <v-expansion-panel>
          <v-expansion-panel-header>ダイス</v-expansion-panel-header>
          <v-expansion-panel-content>
            <FourCard :headline="headline.ability" :data="datas.ability"/>
            <ThreeCard :headline="headline.career" :data="datas.career"/>
            <ThreeCard :headline="headline.state" :data="datas.state"/>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
  </div>
  <!-- /.panel-wrapper -->
</template>

<script>
import ThreeCard from "~/components/ui/ThreeCard";
import FourCard from "~/components/ui/FourCard";

export default {
  props: ["datas"],
  components: {
    ThreeCard,
    FourCard
  },
  data() {
    return {
      headline: {
        ability: "能力値",
        career: "経歴",
        state: "状態"
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.panel-wrapper {
  margin: 30px auto;
}
</style>