<template>
  <div class="container">
    <v-app>
    <SideButton />
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
      <!-- 送るデータは後で変更 -->
      <ArmorPanel :datas="current.base" />
    </div>
    <!-- /.content-wrapper -->
    </v-app>
    <!-- {{ getDice }} -->
  </div>
</template>

<script>
import ThreeCard from "~/components/ui/ThreeCard";
import SideButton from "~/components/ui/SideButton";
import DicePanel from "~/components/layouts/DicePanel";
import BasePanel from "~/components/layouts/BasePanel";
import ArmorPanel from "~/components/layouts/ArmorPanel";
import datas from "~/assets/data/users.json";
import dice from "~/plugins/dice";
import mapGetters from "vuex";

export default {
  components: {
    ThreeCard,
    DicePanel,
    BasePanel,
    ArmorPanel,
    SideButton
  },
  data() {
    return {
      datas,
      current: {
        name: "Taro",
        dice: {
          ability: {
            name: "能力値",
            value: {
              stamina: { name: "体力点", value: 0 },
              soul: { name: "魂魄点", value: 0 },
              quantity: { name: "技量点", value: 0 },
              intellectual: { name: "知力点", value: 0 },
              concentrate: { name: "集中点", value: 0 },
              enduring: { name: "持久点", value: 0 },
              reflections: { name: "反射点", value: 0 }
            }
          },
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
      this.current = data;
      this.$store.dispatch("user/userChange", data);
      console.log(this.$store.state.user);
    }
  },
  computed: {
    // ...mapGetters("user", ["getDice"])
  }
}
</script>

<style lang="scss" scoped>
.content-wrapper {
  width: 80%;
  max-width: 1000px;
  min-height: 100vh;
  border-radius: 10px;
  position: relative;
  margin: 0 auto;
  top: -150px;
}
</style>