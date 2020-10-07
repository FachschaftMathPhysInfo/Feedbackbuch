<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        <v-toolbar-title
          ><v-icon class="mr-2">mdi-comment-quote</v-icon
          >Feedbackbuch</v-toolbar-title
        >
      </div>

      <v-spacer></v-spacer>

      <span>
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <v-btn v-bind="attrs" v-on="on" v-on:click="yesterday()" icon
              ><v-icon>mdi-menu-left</v-icon>
            </v-btn>
          </template>
          <span>zu vorherigem Tag wechseln</span>
        </v-tooltip>
        Seite vom {{ day | dateString }}
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <v-btn v-bind="attrs" v-on="on" v-on:click="tomorrow()" icon
              ><v-icon>mdi-menu-right</v-icon>
            </v-btn>
          </template>
          <span>zu nächstem Tag wechseln</span>
        </v-tooltip>
      </span>

      <v-spacer></v-spacer>
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            v-bind="attrs"
            v-on="on"
            text
            v-on:click="$vuetify.theme.dark = !$vuetify.theme.dark"
          >
            <v-icon>mdi-theme-light-dark</v-icon>
          </v-btn>
        </template>
        <span>Darkmode umschalten</span>
      </v-tooltip>
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            v-bind="attrs"
            v-on="on"
            href="https://github.com/TomTomRixRix/Feedbackbuch"
            target="_blank"
            text
          >
            <v-icon>mdi-github</v-icon>
          </v-btn>
        </template>
        <span>Projekt auf GitHub</span>
      </v-tooltip>
    </v-app-bar>

    <v-main>
      <div v-if="!$apollo.loading" style="padding-bottom: 64px;">
        <v-list
          color="secondary"
          v-bind:key="comment.id"
          v-for="comment in comments"
        >
          <Comment :comment="comment" @reply="reply" />
        </v-list>
      </div>
      <div v-if="$apollo.loading">loading ....</div>

      <v-expansion-panels
        v-model="panelOpened"
        style="position: fixed; bottom: 0px; width: 100vw"
      >
        <v-expansion-panel>
          <v-expansion-panel-header color="primary">
            <span class="text-h6" style="color: #f4f1ea;"
              >Neues Feedback hinzufügen</span
            >
            <template v-slot:actions>
              <v-icon style="color: #f4f1ea;">
                mdi-menu-up
              </v-icon>
            </template>
          </v-expansion-panel-header>
          <v-expansion-panel-content color="secondary">
            <v-row>
              <v-col cols="12" sm="4">
                <span>
                  <v-chip
                    v-if="currentReference"
                    class="ma-2"
                    close
                    @click:close="currentReference = null"
                  >
                    Bezieht sich auf Kommentar {{ this.currentReference }}
                  </v-chip>
                </span>
              </v-col>
              <v-col cols="12" sm="4">
                <v-tabs v-model="tab" centered background-color="secondary">
                  <v-tabs-slider color="primary"></v-tabs-slider>
                  <v-tab
                    :style="
                      $vuetify.theme.dark
                        ? 'color:white;'
                        : 'color:rgba(0,0,0,0.54);'
                    "
                    v-for="item in items"
                    :key="item"
                  >
                    {{ item }}
                  </v-tab>
                </v-tabs>
              </v-col>
              <v-col cols="12" sm="4">
                <v-row justify="end" align="center" style="height:100%;">
                  <v-btn
                    v-on:click="senden"
                    text
                    depressed
                    :color="$vuetify.theme.dark ? 'white' : 'primary'"
                    style="margin-right: 10px"
                  >
                    Senden
                    <v-icon class="ml-2">mdi-send</v-icon>
                  </v-btn>
                </v-row>
              </v-col>
            </v-row>

            <v-tabs-items v-model="tab" style="background-color:transparent;">
              <v-tab-item v-for="item in items" :key="item">
                <div v-if="item == 'vorschau'" style="padding: 16px">
                  <v-card>
                    <Editor
                      mode="viewer"
                      hint="Preview"
                      :emoji="true"
                      :image="false"
                      :outline="true"
                      :render-config="renderConfig"
                      v-model="text"
                    />
                  </v-card>
                </div>
                <div v-if="item == 'editor'" style="margin-top: -16px">
                  <Editor
                    :outline="false"
                    mode="editor"
                    hint="Hint: Use '$$ \LaTeX $$' to write math formulas. Place an empty line before your formula to end markdown code."
                    :emoji="true"
                    :image="false"
                    :render-config="renderConfig"
                    v-model="text"
                    counter="31415"
                  />
                </div>
              </v-tab-item>
            </v-tabs-items>
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
import gql from "graphql-tag";

const COMMENTS_QUERY = gql`
  query {
    comments {
      timestamp
      content
      id
      upvotes
      references
    }
  }
`;

export default {
  name: "App",

  components: {
    Comment,
    // Day,
    Editor,
  },

  data() {
    return {
      panelOpened: 1,
      currentReference: null,
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
  computed: {},
  methods: {
    senden() {
      this.$apollo.mutate({
        mutation: gql`
          mutation($content: String!, $references: Int) {
            createComment(content: $content, references: $references) {
              id
              content
              upvotes
              timestamp
              references
            }
          }
        `,
        variables: {
          content: this.text,
          references: this.currentReference,
        },
      });

      //reset UI for next comment
      this.text = "";
      this.currentReference = null;
      this.panelOpened = 1; //everything else than 0 works
    },
    yesterday() {
      this.daysOffsetCounter = this.daysOffsetCounter - 1;
      this.day = moment(this.today).add(this.daysOffsetCounter, "days"); //.format('LL');
    },
    tomorrow() {
      this.daysOffsetCounter = this.daysOffsetCounter + 1;
      this.day = moment(this.today).add(this.daysOffsetCounter, "days"); //.format('LL');
    },
    reply(commentid) {
      this.currentReference = commentid;
    },
  },

  apollo: {
    comments: {
      query: COMMENTS_QUERY,
      update: (data) => data.comments,
      subscribeToMore: {
        document: gql`
          subscription name {
            commentAdded {
              id
              content
              timestamp
              upvotes
              references
            }
          }
        `,
        updateQuery: (previousResult, { subscriptionData }) => {
          if (previousResult == null) {
            return {
              comments: [subscriptionData.data.commentAdded],
            };
          } else {
            return {
              comments: [
                ...previousResult.comments,
                subscriptionData.data.commentAdded,
              ],
            };
          }
        },
      },
    },
  },

  mounted() {
    // Access properties or methods using $refs
    // this.$refs.editor.focus();
    // this.$refs.editor.upload();
    // Dark theme
    // this.$vuetify.theme.dark = true;
  },
  filters: {
    dateString: function(now) {
      now.locale("de");
      return now.format("LL");
    },
  },
};
</script>

<style scoped></style>
