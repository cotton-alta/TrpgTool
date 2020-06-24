<template>
  <div>
    <v-app>
    <div class="content-wrapper">
      <v-bottom-navigation
        grow
        color="teal"
      >
        <v-btn
          v-for="data in datas"
          :key="data.id"
          @click="changeCharacter(data)"
        >
          <span>{{ data.name }}</span>
        </v-btn>
      </v-bottom-navigation>
      <DicePanel :datas="current.dice" />
      <BasePanel :datas="current.base" />
    </div>
    <!-- /.content-wrapper -->
    <div class="side-wrapper">

    </div>
    <!-- /.side-wrapper -->
    </v-app>
  </div>
</template>

<script>
import ThreeCard from "~/components/ui/ThreeCard";
import DicePanel from "~/components/layouts/DicePanel";
import BasePanel from "~/components/layouts/BasePanel";
import datas from "~/assets/data/users.json";

export default {
  components: {
    ThreeCard,
    DicePanel,
    BasePanel
  },
  data() {
    return {
      datas,
      current: {
        name: "Taro",
        dice: {
          ability: { name: "能力値", value: 0 },
          career: {
            name: "経歴",
            value: {
              origins: { name: "出自", value: 0 },
              history: { name: "来歴", value: 0 },
              chance: { name: "邂逅", value: 0 }
            }
          },
          state: { name: "状態", value: 0 }
        },
        base: {
          ability: { name: "能力値", value: 0 },
          career: { name: "経歴", value: 0 }
        }
      }
    }
  },
  methods: {
    changeCharacter: function(data) {
      this.current = data
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
</style>