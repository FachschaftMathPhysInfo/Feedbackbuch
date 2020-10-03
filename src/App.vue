<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        <v-toolbar-title>Feedbackbuch</v-toolbar-title>
      </div>

      <v-spacer></v-spacer>

      <span
        ><v-btn v-on:click="yesterday()" icon
          ><v-icon>mdi-menu-left</v-icon>
        </v-btn>
        {{ day | dateString }}
        <v-btn v-on:click="tomorrow()" icon
          ><v-icon>mdi-menu-right</v-icon>
        </v-btn></span
      >

      <v-spacer></v-spacer>
      <v-btn
        href="https://github.com/vuetifyjs/vuetify/releases/latest"
        target="_blank"
        text
      >
        <v-icon>mdi-github</v-icon>
      </v-btn>
    </v-app-bar>

    <v-main>
      <v-list v-bind:key="comment.id" v-for="comment in comments">
        <Comment :comment="comment" />
      </v-list>

      <v-expansion-panels
        style="position: absolute; bottom: 10px; width: 100vw;"
      >
        <v-expansion-panel>
          <v-expansion-panel-header expand-icon="mdi-menu-up">
            <h3>Neues Feedback hinzufügen</h3>
          </v-expansion-panel-header>
          <v-expansion-panel-content>
            <v-tabs v-model="tab" right>
              <v-tabs-slider></v-tabs-slider>
              <v-tab v-for="item in items" :key="item">
                {{ item }}
              </v-tab>
            </v-tabs>

            <v-tabs-items v-model="tab" style="min-height: 275px;">
              <v-tab-item v-for="item in items" :key="item">
                <div v-if="item == 'vorschau'">
                  <Editor
                    mode="viewer"
                    hint="Preview"
                    :emoji="true"
                    :image="false"
                    :outline="false"
                    :render-config="renderConfig"
                    v-model="text"
                  />
                </div>
                <div v-if="item == 'editor'">
                  <Editor
                    mode="editor"
                    hint="Hint: Use '$$ \LaTeX $$' to write math formulas. Place an empty line before your formula to end markdown code."
                    :emoji="true"
                    :image="false"
                    :outline="false"
                    :render-config="renderConfig"
                    v-model="text"
                    counter="31415"
                  />
                </div>
              </v-tab-item>
            </v-tabs-items>
            <v-row justify="end">
              <v-btn
                v-on:click="senden"
                text
                depressed
                color="primary"
                style="margin-right:10px"
              >
                Senden
                <v-icon class="ml-2">mdi-send</v-icon>
              </v-btn>
            </v-row>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-main>
  </v-app>
</template>

<script>
import Comment from "./components/Comment";
// import Day from "./components/Day";
import { Editor } from "vuetify-markdown-editor";
import moment from "moment";

export default {
  name: "App",

  components: {
    Comment,
    // Day,
    Editor,
  },

  data() {
    return {
      day: moment(new Date()),
      today: moment(new Date()),
      daysOffsetCounter: 0,
      tab: null,
      items: ["editor", "vorschau"],
      comments: [],
      text: "",
      renderConfig: {
        // Mermaid config
        mermaid: {
          theme: "dark",
        },
        texmath: {
          engine: require("katex"),
          katexOptions: {
            macros: { "\\RR": "\\mathbb{R}" },
            output: "mathml", // formula delimiters
            delimiters: [
              {
                left: "$$$",
                right: "$$$",
                options: {
                  displayMode: true, // block
                },
              },
              {
                left: "\\[",
                right: "\\]",
                options: {
                  displayMode: true, // block
                },
              },
              {
                left: "$$",
                right: "$$",
                options: {
                  displayMode: false, // inline
                },
              },
              {
                left: "\\(",
                right: "\\)",
                options: {
                  displayMode: false, // inline
                },
              },
            ],
          },
        },
      },
    };
  },
  computed: {
    
  },
  methods: {
    senden() {
      this.comments.push({
        text: this.text,
        timestamp: new Date(),
        upvotes: 0,
        id: 0,
      });
    },
    yesterday() {
      this.daysOffsetCounter = this.daysOffsetCounter -1;
      this.day = moment(this.today).add(this.daysOffsetCounter,"days"); //.format('LL');
    },
    tomorrow() {
      this.daysOffsetCounter = this.daysOffsetCounter +1;
      this.day = moment(this.today).add(this.daysOffsetCounter,"days"); //.format('LL');
    },
    
  },
  mounted() {
    // Access properties or methods using $refs
    // this.$refs.editor.focus();
    // this.$refs.editor.upload();

    // Dark theme
    // this.$vuetify.theme.dark = true;
  },
  filters:{
    dateString: function (now){
      now.locale("de");  
      return now.format('LL');
    }
  }
};
</script>
