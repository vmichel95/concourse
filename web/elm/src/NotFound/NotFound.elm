module NotFound.NotFound exposing
    ( documentTitle
    , handleDelivery
    , init
    , subscriptions
    , view
    )

import EffectTransformer exposing (ET)
import Html exposing (Html)
import Html.Attributes exposing (class, href, id, src)
import Login.Login as Login
import Message.Effects exposing (Effect(..))
import Message.Message exposing (Message(..))
import Message.Subscription
    exposing
        ( Delivery(..)
        , Interval(..)
        , Subscription(..)
        )
import Message.TopLevelMessage exposing (TopLevelMessage(..))
import NotFound.Model exposing (Model)
import Routes
import Session exposing (Session)
import SideBar.SideBar as SideBar
import Views.Styles
import Views.TopBar as TopBar


type alias Flags =
    { route : Routes.Route
    , notFoundImgSrc : String
    }


init : Flags -> ( Model, List Effect )
init flags =
    ( { notFoundImgSrc = flags.notFoundImgSrc
      , route = flags.route
      , isUserMenuExpanded = False
      }
    , []
    )


documentTitle : String
documentTitle =
    "Not Found"


view : Session a -> Model -> Html Message
view session model =
    Html.div
        (id "page-including-top-bar" :: Views.Styles.pageIncludingTopBar)
        [ Html.div
            (id "top-bar-app" :: Views.Styles.topBar False)
            [ SideBar.hamburgerMenu session
            , TopBar.concourseLogo
            , TopBar.breadcrumbs model.route
            , Login.view session.userState model False
            ]
        , Html.div
            (id "page-below-top-bar" :: Views.Styles.pageBelowTopBar model.route)
            [ SideBar.view
                { expandedTeams = session.expandedTeams
                , pipelines = session.pipelines
                , hovered = session.hovered
                , isSideBarOpen = session.isSideBarOpen
                , screenSize = session.screenSize
                }
                Nothing
            , Html.div [ class "notfound" ]
                [ Html.div [ class "title" ] [ Html.text "404" ]
                , Html.div [ class "reason" ] [ Html.text "this page was not found" ]
                , Html.img [ src model.notFoundImgSrc ] []
                , Html.div [ class "help-message" ]
                    [ Html.text "Not to worry, you can head"
                    , Html.br [] []
                    , Html.text "back to the "
                    , Html.a [ href "/" ] [ Html.text "home page" ]
                    ]
                ]
            ]
        ]


subscriptions : List Subscription
subscriptions =
    [ OnClockTick FiveSeconds ]


handleDelivery : Delivery -> ET Model
handleDelivery delivery ( model, effects ) =
    case delivery of
        ClockTicked FiveSeconds _ ->
            ( model, effects ++ [ FetchPipelines ] )

        _ ->
            ( model, effects )
