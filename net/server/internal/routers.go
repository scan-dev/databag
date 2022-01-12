/*
 * DataBag
 *
 * DataBag provides storage for decentralized identity based self-hosting apps. It is intended to support sharing of personal data and hosting group conversations. 
 *
 * API version: 0.0.1
 * Contact: roland.osborne@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-databag/swagger-codegen.git)
 */
package databag

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

  // populate context
  _configured = getBoolConfigValue(CONFIG_CONFIGURED, false);
  _adminUsername = getStrConfigValue(CONFIG_USERNAME, "");
  _adminPassword = getBinConfigValue(CONFIG_PASSWORD, nil);
  _nodeDomain = getStrConfigValue(CONFIG_DOMAIN, "");
  _publicLimit = getNumConfigValue(CONFIG_PUBLICLIMIT, 0);
  _accountStorage = getNumConfigValue(CONFIG_STORAGE, 0);

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AddAccount",
		strings.ToUpper("Post"),
		"/account/profile",
		AddAccount,
	},

	Route{
		"AddAccountApp",
		strings.ToUpper("Post"),
		"/account/apps",
		AddAccountApp,
	},

	Route{
		"AddAccountAuthentication",
		strings.ToUpper("Post"),
		"/account/auth",
		AddAccountAuthentication,
	},

	Route{
		"AddPublicAccount",
		strings.ToUpper("Post"),
		"/account/public/profile",
		AddPublicAccount,
	},

	Route{
		"GetAccountApps",
		strings.ToUpper("Get"),
		"/account/apps",
		GetAccountApps,
	},

	Route{
		"GetAccountImage",
		strings.ToUpper("Get"),
		"/account/profile/image",
		GetAccountImage,
	},

	Route{
		"GetAccountProfile",
		strings.ToUpper("Get"),
		"/account/profile",
		GetAccountProfile,
	},

	Route{
		"GetAccountToken",
		strings.ToUpper("Get"),
		"/account/token",
		GetAccountToken,
	},

	Route{
		"GetAccountUsername",
		strings.ToUpper("Get"),
		"/account/claimable",
		GetAccountUsername,
	},

	Route{
		"GetPublicClaimable",
		strings.ToUpper("Get"),
		"/account/public/claimable",
		GetPublicClaimable,
	},

	Route{
		"RemoveAccountApp",
		strings.ToUpper("Delete"),
		"/account/apps/{appId}",
		RemoveAccountApp,
	},

	Route{
		"SetAccountApp",
		strings.ToUpper("Put"),
		"/account/apps",
		SetAccountApp,
	},

	Route{
		"SetAccountAuthentication",
		strings.ToUpper("Put"),
		"/account/auth",
		SetAccountAuthentication,
	},

	Route{
		"AddNodeAccount",
		strings.ToUpper("Post"),
		"/admin/accounts",
		AddNodeAccount,
	},

	Route{
		"GetNodeAccountImage",
		strings.ToUpper("Get"),
		"/admin/accounts/{accountId}/image",
		GetNodeAccountImage,
	},

	Route{
		"GetNodeAccounts",
		strings.ToUpper("Get"),
		"/admin/accounts",
		GetNodeAccounts,
	},

	Route{
		"GetNodeClaimable",
		strings.ToUpper("Get"),
		"/admin/claimable",
		GetNodeClaimable,
	},

	Route{
		"GetNodeConfig",
		strings.ToUpper("Get"),
		"/admin/config",
		GetNodeConfig,
	},

	Route{
		"RemoveNodeAccount",
		strings.ToUpper("Delete"),
		"/admin/accounts/{accountId}",
		RemoveNodeAccount,
	},

	Route{
		"SetNodeAccount",
		strings.ToUpper("Put"),
		"/admin/accounts/{accountId}/reset",
		SetNodeAccount,
	},

	Route{
		"SetNodeClaim",
		strings.ToUpper("Post"),
		"/admin/claim",
		SetNodeClaim,
	},

	Route{
		"SetNodeConfig",
		strings.ToUpper("Put"),
		"/admin/config",
		SetNodeConfig,
	},

	Route{
		"Authorize",
		strings.ToUpper("Put"),
		"/authorize",
		Authorize,
	},

	Route{
		"AddCard",
		strings.ToUpper("Post"),
		"/contact/cards",
		AddCard,
	},

	Route{
		"ClearCardGroup",
		strings.ToUpper("Delete"),
		"/contact/cards/{cardId}/groups/{groupId}",
		ClearCardGroup,
	},

	Route{
		"ClearCardNotes",
		strings.ToUpper("Delete"),
		"/contact/cards/{cardId}/notes",
		ClearCardNotes,
	},

	Route{
		"GetCard",
		strings.ToUpper("Get"),
		"/contact/cards/{cardId}",
		GetCard,
	},

	Route{
		"GetCardData",
		strings.ToUpper("Get"),
		"/contact/cards/{cardId}/data",
		GetCardData,
	},

	Route{
		"GetCardProfile",
		strings.ToUpper("Get"),
		"/contact/cards/{cardId}/profile",
		GetCardProfile,
	},

	Route{
		"GetCardProfileImage",
		strings.ToUpper("Get"),
		"/contact/cards/{cardId}/profile/image",
		GetCardProfileImage,
	},

	Route{
		"GetCardView",
		strings.ToUpper("Get"),
		"/contact/cards/view",
		GetCardView,
	},

	Route{
		"GetCloseMessage",
		strings.ToUpper("Get"),
		"/contact/cards/{cardId}/closeMessage",
		GetCloseMessage,
	},

	Route{
		"GetOpenMessage",
		strings.ToUpper("Get"),
		"/contact/cards/{cardId}/openMessage",
		GetOpenMessage,
	},

	Route{
		"RemoveCard",
		strings.ToUpper("Delete"),
		"/contact/cards/{cardId}",
		RemoveCard,
	},

	Route{
		"SetCardGroup",
		strings.ToUpper("Put"),
		"/contact/cards/{cardId}/groups/{groupId}",
		SetCardGroup,
	},

	Route{
		"SetCardNotes",
		strings.ToUpper("Put"),
		"/contact/cards/{cardId}/notes",
		SetCardNotes,
	},

	Route{
		"SetCardProfile",
		strings.ToUpper("Put"),
		"/contact/cards/{cardId}/profile",
		SetCardProfile,
	},

	Route{
		"SetCardStatus",
		strings.ToUpper("Put"),
		"/contact/cards/{cardId}/status",
		SetCardStatus,
	},

	Route{
		"SetCloseMessage",
		strings.ToUpper("Put"),
		"/contact/closeMessage",
		SetCloseMessage,
	},

	Route{
		"SetContentRevision",
		strings.ToUpper("Put"),
		"/contact/content/revision",
		SetContentRevision,
	},

	Route{
		"SetOpenMessage",
		strings.ToUpper("Put"),
		"/contact/openMessage",
		SetOpenMessage,
	},

	Route{
		"SetProfileRevision",
		strings.ToUpper("Put"),
		"/contact/profile/revision",
		SetProfileRevision,
	},

	Route{
		"AddArticle",
		strings.ToUpper("Post"),
		"/content/articles",
		AddArticle,
	},

	Route{
		"AddArticleAsset",
		strings.ToUpper("Post"),
		"/content/articles/{articleId}/assets",
		AddArticleAsset,
	},

	Route{
		"AddArticleTag",
		strings.ToUpper("Post"),
		"/content/articles/{articleId}/tags",
		AddArticleTag,
	},

	Route{
		"AddLabel",
		strings.ToUpper("Post"),
		"/content/labels",
		AddLabel,
	},

	Route{
		"ClearArticleGroup",
		strings.ToUpper("Delete"),
		"/content/articles/{articleId}/groups/{groupId}",
		ClearArticleGroup,
	},

	Route{
		"ClearArticleLabel",
		strings.ToUpper("Delete"),
		"/content/articles/{articleId}/labels/{labelId}",
		ClearArticleLabel,
	},

	Route{
		"ClearLabelGroup",
		strings.ToUpper("Delete"),
		"/content/labels/{labelId}/groups/{groupId}",
		ClearLabelGroup,
	},

	Route{
		"GetArticle",
		strings.ToUpper("Get"),
		"/content/articles/{articleId}",
		GetArticle,
	},

	Route{
		"GetArticleAsset",
		strings.ToUpper("Get"),
		"/content/articles/{articleId}/assets/{assetId}",
		GetArticleAsset,
	},

	Route{
		"GetArticleAssets",
		strings.ToUpper("Get"),
		"/content/articles/{articleId}/assets",
		GetArticleAssets,
	},

	Route{
		"GetArticleBlockView",
		strings.ToUpper("Get"),
		"/content/articleBlocks/view",
		GetArticleBlockView,
	},

	Route{
		"GetArticleSubjectField",
		strings.ToUpper("Get"),
		"/content/articles/{articleId}/subject/{field}",
		GetArticleSubjectField,
	},

	Route{
		"GetArticleTag",
		strings.ToUpper("Get"),
		"/content/articles/{articleId}/tags/{tagId}",
		GetArticleTag,
	},

	Route{
		"GetArticleTagBlockView",
		strings.ToUpper("Get"),
		"/content/articles/{articleId}/tagBlocks/view",
		GetArticleTagBlockView,
	},

	Route{
		"GetArticleTagSubjectField",
		strings.ToUpper("Get"),
		"/content/articles/{articleId}/tags/{tagId}/subject/{field}",
		GetArticleTagSubjectField,
	},

	Route{
		"GetArticleTagView",
		strings.ToUpper("Get"),
		"/content/articles/{articleId}/tagBlocks/{blockId}/view",
		GetArticleTagView,
	},

	Route{
		"GetArticleTags",
		strings.ToUpper("Get"),
		"/content/articles/{articleId}/tagBlocks/{blockId}",
		GetArticleTags,
	},

	Route{
		"GetArticleViews",
		strings.ToUpper("Get"),
		"/content/articleBlocks/{blockId}/view",
		GetArticleViews,
	},

	Route{
		"GetArticles",
		strings.ToUpper("Get"),
		"/content/articleBlocks/{blockId}",
		GetArticles,
	},

	Route{
		"GetLabels",
		strings.ToUpper("Get"),
		"/content/labels",
		GetLabels,
	},

	Route{
		"RemoveArticle",
		strings.ToUpper("Delete"),
		"/content/articles/{articleId}",
		RemoveArticle,
	},

	Route{
		"RemoveArticleAsset",
		strings.ToUpper("Delete"),
		"/content/articles/{articleId}/assets/{assetId}",
		RemoveArticleAsset,
	},

	Route{
		"RemoveArticleTag",
		strings.ToUpper("Delete"),
		"/content/articles/{articleId}/tags/{tagId}",
		RemoveArticleTag,
	},

	Route{
		"RemoveLabel",
		strings.ToUpper("Delete"),
		"/content/labels/{labelId}",
		RemoveLabel,
	},

	Route{
		"SetArticleConfirmed",
		strings.ToUpper("Put"),
		"/content/articles/{articleId}/confirmed",
		SetArticleConfirmed,
	},

	Route{
		"SetArticleGroup",
		strings.ToUpper("Post"),
		"/content/articles/{articleId}/groups/{groupId}",
		SetArticleGroup,
	},

	Route{
		"SetArticleLabel",
		strings.ToUpper("Post"),
		"/content/articles/{articleId}/labels/{labelId}",
		SetArticleLabel,
	},

	Route{
		"SetArticleSubject",
		strings.ToUpper("Put"),
		"/content/articles/{articleId}/subject",
		SetArticleSubject,
	},

	Route{
		"SetLabelGroup",
		strings.ToUpper("Post"),
		"/content/labels/{labelId}/groups/{groupId}",
		SetLabelGroup,
	},

	Route{
		"UpdateLabel",
		strings.ToUpper("Put"),
		"/content/labels/{labelId}",
		UpdateLabel,
	},

	Route{
		"AddDialogue",
		strings.ToUpper("Post"),
		"/conversation/dialogues",
		AddDialogue,
	},

	Route{
		"AddDialogueInsight",
		strings.ToUpper("Put"),
		"/conversation/dialogues/{dialogueId}/cards/{cardId}",
		AddDialogueInsight,
	},

	Route{
		"AddDialogueTopic",
		strings.ToUpper("Post"),
		"/conversation/dialogues/{dialogueId}/topics",
		AddDialogueTopic,
	},

	Route{
		"AddInsightDialogue",
		strings.ToUpper("Post"),
		"/conversation/insights/{dialogueId}",
		AddInsightDialogue,
	},

	Route{
		"AddTopicAsset",
		strings.ToUpper("Post"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/assets",
		AddTopicAsset,
	},

	Route{
		"AddTopicTag",
		strings.ToUpper("Post"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/tags",
		AddTopicTag,
	},

	Route{
		"ConversationDialoguesDialogueIdTopicsTopicIdConfirmedPut",
		strings.ToUpper("Put"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/confirmed",
		ConversationDialoguesDialogueIdTopicsTopicIdConfirmedPut,
	},

	Route{
		"GetDialogueTopic",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}",
		GetDialogueTopic,
	},

	Route{
		"GetDialogueTopicSubjectField",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/subject/{field}",
		GetDialogueTopicSubjectField,
	},

	Route{
		"GetDialogues",
		strings.ToUpper("Get"),
		"/conversation/dialogues",
		GetDialogues,
	},

	Route{
		"GetInsights",
		strings.ToUpper("Get"),
		"/conversation/insights",
		GetInsights,
	},

	Route{
		"GetTopicAsset",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/assets/{assetId}",
		GetTopicAsset,
	},

	Route{
		"GetTopicAssets",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/assets",
		GetTopicAssets,
	},

	Route{
		"GetTopicBlock",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topicBlocks/{blockId}",
		GetTopicBlock,
	},

	Route{
		"GetTopicBlockView",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topicBlocks/view",
		GetTopicBlockView,
	},

	Route{
		"GetTopicTag",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/tags/{tagId}",
		GetTopicTag,
	},

	Route{
		"GetTopicTagBlockView",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/tagBlocks/view",
		GetTopicTagBlockView,
	},

	Route{
		"GetTopicTagSubjectField",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/tags/{tagId}/subject/{field}",
		GetTopicTagSubjectField,
	},

	Route{
		"GetTopicTagView",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/tagBlocks/{blockId}/view",
		GetTopicTagView,
	},

	Route{
		"GetTopicTags",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/tagBlocks/{blockId}",
		GetTopicTags,
	},

	Route{
		"GetTopicViews",
		strings.ToUpper("Get"),
		"/conversation/dialogues/{dialogueId}/topicBlocks/{blockId}/view",
		GetTopicViews,
	},

	Route{
		"RemoveDialogue",
		strings.ToUpper("Delete"),
		"/conversation/dialogues/{dialogueId}",
		RemoveDialogue,
	},

	Route{
		"RemoveDialogueInsight",
		strings.ToUpper("Delete"),
		"/conversation/dialogues/{dialogueId}/cards/{cardId}",
		RemoveDialogueInsight,
	},

	Route{
		"RemoveDialogueTopic",
		strings.ToUpper("Delete"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}",
		RemoveDialogueTopic,
	},

	Route{
		"RemoveTopicAsset",
		strings.ToUpper("Delete"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/assets/{assetId}",
		RemoveTopicAsset,
	},

	Route{
		"RemoveTopicTag",
		strings.ToUpper("Delete"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/tags/{tagId}",
		RemoveTopicTag,
	},

	Route{
		"SetDialogueActive",
		strings.ToUpper("Put"),
		"/conversation/dialogues/{dialogueId}/active",
		SetDialogueActive,
	},

	Route{
		"SetDialogueInsightStatus",
		strings.ToUpper("Put"),
		"/conversation/dialogues/{dialogueId}/status",
		SetDialogueInsightStatus,
	},

	Route{
		"SetDialogueSubject",
		strings.ToUpper("Put"),
		"/conversation/dialogues/{dialogueId}/subject",
		SetDialogueSubject,
	},

	Route{
		"SetInsightDialogue",
		strings.ToUpper("Delete"),
		"/conversation/insights/{dialogueId}",
		SetInsightDialogue,
	},

	Route{
		"SetInsightStatus",
		strings.ToUpper("Put"),
		"/conversation/insights/{insightId}/status",
		SetInsightStatus,
	},

	Route{
		"SetTopicSubject",
		strings.ToUpper("Put"),
		"/conversation/dialogues/{dialogueId}/topics/{topicId}/subject",
		SetTopicSubject,
	},

	Route{
		"GetProfile",
		strings.ToUpper("Get"),
		"/profile",
		GetProfile,
	},

	Route{
		"GetProfileImage",
		strings.ToUpper("Get"),
		"/profile/image",
		GetProfileImage,
	},

	Route{
		"GetProfileMessage",
		strings.ToUpper("Get"),
		"/profile/message",
		GetProfileMessage,
	},

	Route{
		"SetProfile",
		strings.ToUpper("Put"),
		"/profile",
		SetProfile,
	},

	Route{
		"AddGroup",
		strings.ToUpper("Post"),
		"/share/groups",
		AddGroup,
	},

	Route{
		"GetGroups",
		strings.ToUpper("Get"),
		"/share/groups",
		GetGroups,
	},

	Route{
		"RemoveGroup",
		strings.ToUpper("Delete"),
		"/share/groups/{groupId}",
		RemoveGroup,
	},

	Route{
		"UpdateGroup",
		strings.ToUpper("Put"),
		"/share/groups/{groupId}",
		UpdateGroup,
	},

	Route{
		"Status",
		strings.ToUpper("Get"),
		"/status",
		Status,
	},
}
