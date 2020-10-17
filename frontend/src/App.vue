<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        <v-toolbar-title
          ><v-icon
            v-on:click="login"
            class="mr-2"
            :style="this.admin ? 'color:green;' : 'color:white;'"
            >mdi-comment-quote</v-icon
          >Feedbackbuch</v-toolbar-title
        >
      </div>

      <v-spacer></v-spacer>

      <span>
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <v-btn :disabled="disableYesterday"  v-bind="attrs" v-on="on" v-on:click="yesterday()" icon
              ><v-icon>mdi-menu-left</v-icon>
            </v-btn>
          </template>
          <span>zu vorherigem Tag wechseln</span>
        </v-tooltip>
        Seite vom {{ day | dateString }}
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <v-btn :disabled="disableTomorrow" v-bind="attrs" v-on="on" v-on:click="tomorrow()" icon
              ><v-icon>mdi-menu-right</v-icon>
            </v-btn>
          </template>
          <span>zu nächstem Tag wechseln</span>
        </v-tooltip>
      </span>

      <v-spacer></v-spacer>
      <v-btn-toggle borderless dense light v-model="sorting">
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              value="time"
              v-bind="attrs"
              v-on="on"
              v-on:click="sortByTime"
            >
              <v-icon left class="ml-2">
                mdi-sort-clock-ascending-outline
              </v-icon>
              <span class="hidden-sm-and-down">Time</span>
            </v-btn>
          </template>
          <span>Chronologisch sortieren</span>
        </v-tooltip>
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              value="votes"
              v-bind="attrs"
              v-on="on"
              v-on:click="sortByUpvotes"
            >
              <v-icon left class="ml-2"> mdi-sort-bool-ascending </v-icon>
              <span class="hidden-sm-and-down">Votes</span>
            </v-btn>
          </template>
          <span>nach Votes sortieren</span>
        </v-tooltip>
      </v-btn-toggle>
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            v-bind="attrs"
            v-on="on"
            text
            href="https://mathphys.info/vorkurs/plan/"
            target="_blank"
          >
            <v-icon>mdi-calendar</v-icon>
          </v-btn>
        </template>
        <span>Vorkursplan</span>
      </v-tooltip>
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

    <v-main class="secondary">
      <div v-if="!$apollo.loading" >
        <v-card v-if="todays(comments).length === 0" elevation="12" style="margin: 12px; padding: 12px; text-align:center">
          <h2>Es liegt (noch) kein Feedback für diesen Tag vor.<br>
        Wir würden uns aber über dein Feedback freuen 😀</h2></v-card> 
        <v-list
          color="secondary"
          v-bind:key="comment.id"
          v-for="comment in todays(comments)"
        >
          <Comment :id="'comment-'+comment.id" :highlighted="highlightedCommentID === comment.id" :comment="comment" :admin="admin" @reply="reply" @jumpToComment="jumpToComment"/>
        </v-list>
      </div>
      <div v-if="$apollo.loading">loading ....</div>

      <v-expansion-panels v-model="panelOpened" style="position: sticky; bottom: 0px; width: 100vw">
        <v-expansion-panel>
          <v-expansion-panel-header color="primary">
            <span class="text-h6" style="color: #f4f1ea"
              >Neues Feedback hinzufügen</span
            >
            <template v-slot:actions>
              <v-icon style="color: #f4f1ea"> mdi-menu-up </v-icon>
            </template>
          </v-expansion-panel-header>
          <v-expansion-panel-content color="secondary">
            <v-row align="center">
              <v-col cols="12" sm="4">
                <v-tooltip right>
                  <template v-slot:activator="{ on, attrs }">
                    <span v-bind="attrs" v-on="on"><v-icon>mdi-information</v-icon> Nettiquette</span>
                  </template>
                  <span>Bitte bedenke, dass das Feedback anonym, jedoch trotzdem öffentlich einsehbar ist. Überlege dir also gut, wie du dein Feedback formulierst. Bitte sei respektvoll und beleidige niemanden, ein toleranter Umgangston und konstruktive Kritik eingen sich am ehesten, um Veränderungen anzustoßen.</span>
                </v-tooltip>
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
            </v-row>
            <v-tabs-items v-model="tab" style="background-color: transparent">
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
            <v-row justify="space-between">
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
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-main>
  </v-app>
</template>

<script>
import Comment from "./components/Comment";
import { Editor } from "vuetify-markdown-editor";
import moment from "moment";
import gql from "graphql-tag";
import sha256 from "crypto-js/sha256";
import Base64 from "crypto-js/enc-base64";

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
    Editor,
  },

  data() {
    return {
      highlightedCommentID: 0,
      disableTomorrow: true,
      panelOpened: 1,
      admin: false,
      tribleClickCounter: 0,
      sorting: "time",
      currentReference: null,
      day: moment(new Date()),
      today: moment(new Date()),
      daysOffsetCounter: 0,
      tab: null,
      items: ["editor", "vorschau"],
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
    todays(comments){
      //man soll nicht in die Zukunft blättern können
      this.disableTomorrow = !moment(this.day).isBefore(this.today);

      //man soll nicht weiter zurück als der erste Kommentar können
      this.disableYesterday = this.comments.every(comment => moment(comment.timestamp).isAfter(moment(this.day).subtract(1,'days'), 'day'));
      
      //zeige nur die Kommentare des aktuell ausgewählten Tages
      return comments.filter(comment => moment(comment.timestamp).isSame(this.day, 'day'));
    },
    reply(commentid) {
      this.currentReference = commentid;
      this.panelOpened = 0; //everything else than 0 works
    },
    jumpToComment(referenceID){
      if(this.comments){
        this.day = moment(this.comments.filter(comment => comment.id === referenceID)[0].timestamp);
        this.daysOffsetCounter = this.day.diff(this.today,'days');
        if( document.getElementById("comment-"+referenceID)){
          document.getElementById("comment-"+referenceID).scrollIntoView();
        }
        this.highlightedCommentID = referenceID;
      }
      
    },
    sortByTime() {
      this.comments = this.comments.sort((a, b) => {
        return moment(a.timestamp) > moment(b.timestamp);
      });
    },
    sortByUpvotes() {
      this.comments = this.comments.sort((a, b) => {
        return b.upvotes - a.upvotes;
      });
    },
    login() {
      this.tribleClickCounter = this.tribleClickCounter + 1;
      if (this.tribleClickCounter >= 3) {
        this.admin =
          Base64.stringify(sha256(prompt("Sesam öffne dich..."))) ===
          "FkXyH+oqft+3ZpDNSPnlXIH5Dm8qNiLyC2s5ubL4nq4=";
      }
    },
  },

  apollo: {
    comments: {
      query: COMMENTS_QUERY,
      update: (data) => data.comments,
      subscribeToMore: {
        document: gql`
          subscription name {
            commentChanged {
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
              comments: [subscriptionData.data.commentChanged],
            };
          } else {
            //Comment is in previousResults --> should be deleted
            if (
              previousResult.comments.some(
                (c) => c.id === subscriptionData.data.commentChanged.id
              )
            ) {
              const index = previousResult.comments.findIndex(
                (c) => c.id === subscriptionData.data.commentChanged.id
              );

              if (index === -1) return previousResult;

              // The previous result is immutable
              const newResult = {
                comments: [...previousResult.comments],
              };
              // Remove the comment from the list
              newResult.comments.splice(index, 1);
              // Remove references to this comment
              newResult.comments.map((c) => {
                if (c.references === subscriptionData.data.commentChanged.id) {
                  c.references = -1;
                }
              });
              return newResult;
            } else {
              //Comment is NOT in previousResults --> should be added
              return {
                comments: [
                  ...previousResult.comments,
                  subscriptionData.data.commentChanged,
                ],
              };
            }
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
    if (
      window.matchMedia &&
      window.matchMedia("(prefers-color-scheme: dark)").matches
    ) {
      // dark mode
      this.$vuetify.theme.dark = true;
    }

    window
      .matchMedia("(prefers-color-scheme: dark)")
      .addEventListener("change", (e) => {
        this.$vuetify.theme.dark = e.matches;
      });
    //
  },
  filters: {
    dateString: function (now) {
      now.locale("de");
      return now.format("LL");
    },
  },
};
</script>

<style scoped></style>
