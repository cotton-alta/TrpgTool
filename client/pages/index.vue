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
import AbilityCard from "~/components/ui/AbilityCard";
import DicePanel from "~/components/layouts/DicePanel";
import BasePanel from "~/components/layouts/BasePanel";

export default {
  components: {
    AbilityCard,
    DicePanel,
    BasePanel
  },
  data() {
    return {
      current: {
        name: "Taro",
        dice: {
          ability: { name: "能力値", value: 0 },
          career: { name: "経歴", value: 0 },
          state: { name: "状態", value: 0 }
        },
        base: {
          ability: { name: "能力値", value: 0 },
          career: { name: "経歴", value: 0 }
        }
      },
      datas: [
        {
          name: "Taro",
          dice: {
            ability: { name: "能力値", value: 1 },
            career: { name: "経歴", value: 0 },
            state: { name: "状態", value: 0 }
          },
          base: {
            ability: { name: "能力値", value: 0 },
            career: { name: "経歴", value: 0, state: 0 }
          }
        },
        {
          name: "Suzuki",
          dice: {
            ability: { name: "能力値", value: 2 },
            career: { name: "経歴", value: 0 },
            state: { name: "状態", value: 0 }
          },
          base: {
            ability: { name: "能力値", value: 0 },
            career: { name: "経歴", value: 0, state: 0 }
          }
        },
        {
          name: "Kato",
          dice: {
            ability: { name: "能力値", value: 3 },
            career: { name: "経歴", value: 0 },
            state: { name: "状態", value: 0 }
          },
          base: {
            ability: { name: "能力値", value: 0 },
            career: { name: "経歴", value: 0, state: 0 }
          }
        },
        {
          name: "Sato",
          dice: {
            ability: { name: "能力値", value: 4 },
            career: { name: "経歴", value: 0 },
            state: { name: "状態", value: 0 }
          },
          base: {
            ability: { name: "能力値", value: 0 },
            career: { name: "経歴", value: 0, state: 0 }
          }
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

// .panel-wrapper {
//   margin: 30px auto;
// }
</style>