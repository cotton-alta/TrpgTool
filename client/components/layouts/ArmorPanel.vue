<template>
  <div class="panel-wrapper">
      <v-expansion-panels
        accordion
        class="panel-wrapper"
      >
        <v-expansion-panel>
          <v-expansion-panel-header>基本能力値</v-expansion-panel-header>
          <v-expansion-panel-content>
            <ThreeCard :headline="headline.first_ability" :data="datas.first_ability"/>
            <ThreeCard :headline="headline.second_ability" :data="datas.second_ability"/>
            <TwoApartCard :data="datas" />
            <TwoApartCard :data="datas" />
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
  </div>
  <!-- /.panel-wrapper -->
</template>

<script>
import ThreeCard from "~/components/ui/ThreeCard";
import TwoApartCard from "~/components/ui/TwoApartCard";

export default {
  props: ["datas"],
  components: {
    ThreeCard,
    TwoApartCard
  },
  data() {
    return {
      headline: {
        first_ability: "第一能力値",
        second_ability: "第二能力値"
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