package packetid

//go:generate stringer -type ClientboundPacketID
//go:generate stringer -type ServerboundPacketID
type (
	ClientboundPacketID int32
	ServerboundPacketID int32
)

// Login Clientbound
const (
	ClientboundLoginLoginDisconnect ClientboundPacketID = iota
	ClientboundLoginHello
	ClientboundLoginGameProfile
	ClientboundLoginLoginCompression
	ClientboundLoginCustomQuery
	ClientboundLoginCookieRequest
)

// Login Serverbound
const (
	ServerboundLoginHello ServerboundPacketID = iota
	ServerboundLoginKey
	ServerboundLoginCustomQueryAnswer
	ServerboundLoginLoginAcknowledged
	ServerboundLoginCookieResponse
)

// Status Clientbound
const (
	ClientboundStatusStatusResponse ClientboundPacketID = iota
	ClientboundStatusPongResponse
)

// Status Serverbound
const (
	ServerboundStatusStatusRequest ServerboundPacketID = iota
	ServerboundStatusPingRequest
)

// Configuration Clientbound
const (
	ClientboundConfigCookieRequest ClientboundPacketID = iota
	ClientboundConfigCustomPayload
	ClientboundConfigDisconnect
	ClientboundConfigFinishConfiguration
	ClientboundConfigKeepAlive
	ClientboundConfigPing
	ClientboundConfigResetChat
	ClientboundConfigRegistryData
	ClientboundConfigResourcePackPop
	ClientboundConfigResourcePackPush
	ClientboundConfigStoreCookie
	ClientboundConfigTransfer
	ClientboundConfigUpdateEnabledFeatures
	ClientboundConfigUpdateTags
	ClientboundConfigSelectKnownPacks
	ClientboundConfigCustomReportDetails
	ClientboundConfigServerLinks
)

// Configuration Serverbound
const (
	ServerboundConfigClientInformation ServerboundPacketID = iota
	ServerboundConfigCookieResponse
	ServerboundConfigCustomPayload
	ServerboundConfigFinishConfiguration
	ServerboundConfigKeepAlive
	ServerboundConfigPong
	ServerboundConfigResourcePack
	ServerboundConfigSelectKnownPacks
)

// Game Clientbound
const (
	BundleDelimiter ClientboundPacketID = iota
	ClientboundAddEntity
	ClientboundAddExperienceOrb
	ClientboundAnimate
	ClientboundAwardStats
	ClientboundBlockChangedAck
	ClientboundBlockDestruction
	ClientboundBlockEntityData
	ClientboundBlockEvent
	ClientboundBlockUpdate
	ClientboundBossEvent
	ClientboundChangeDifficulty
	ClientboundChunkBatchFinished
	ClientboundChunkBatchStart
	ClientboundChunksBiomes
	ClientboundClearTitles
	ClientboundCommandSuggestions
	ClientboundCommands
	ClientboundContainerClose
	ClientboundContainerSetContent
	ClientboundContainerSetData
	ClientboundContainerSetSlot
	ClientboundCookieRequest
	ClientboundCooldown
	ClientboundCustomChatCompletions
	ClientboundCustomPayload
	ClientboundDamageEvent
	ClientboundDebugSample
	ClientboundDeleteChat
	ClientboundDisconnect
	ClientboundDisguisedChat
	ClientboundEntityEvent
	ClientboundExplode
	ClientboundForgetLevelChunk
	ClientboundGameEvent
	ClientboundHorseScreenOpen
	ClientboundHurtAnimation
	ClientboundInitializeBorder
	ClientboundKeepAlive
	ClientboundLevelChunkWithLight
	ClientboundLevelEvent
	ClientboundLevelParticles
	ClientboundLightUpdate
	ClientboundLogin
	ClientboundMapItemData
	ClientboundMerchantOffers
	ClientboundMoveEntityPos
	ClientboundMoveEntityPosRot
	ClientboundMoveEntityRot
	ClientboundMoveVehicle
	ClientboundOpenBook
	ClientboundOpenScreen
	ClientboundOpenSignEditor
	ClientboundPing
	ClientboundPongResponse
	ClientboundPlaceGhostRecipe
	ClientboundPlayerAbilities
	ClientboundPlayerChat
	ClientboundPlayerCombatEnd
	ClientboundPlayerCombatEnter
	ClientboundPlayerCombatKill
	ClientboundPlayerInfoRemove
	ClientboundPlayerInfoUpdate
	ClientboundPlayerLookAt
	ClientboundPlayerPosition
	ClientboundRecipe
	ClientboundRemoveEntities
	ClientboundRemoveMobEffect
	ClientboundResetScore
	ClientboundResourcePackPop
	ClientboundResourcePackPush
	ClientboundRespawn
	ClientboundRotateHead
	ClientboundSectionBlocksUpdate
	ClientboundSelectAdvancementsTab
	ClientboundServerData
	ClientboundSetActionBarText
	ClientboundSetBorderCenter
	ClientboundSetBorderLerpSize
	ClientboundSetBorderSize
	ClientboundSetBorderWarningDelay
	ClientboundSetBorderWarningDistance
	ClientboundSetCamera
	ClientboundSetCarriedItem
	ClientboundSetChunkCacheCenter
	ClientboundSetChunkCacheRadius
	ClientboundSetDefaultSpawnPosition
	ClientboundSetDisplayObjective
	ClientboundSetEntityData
	ClientboundSetEntityLink
	ClientboundSetEntityMotion
	ClientboundSetEquipment
	ClientboundSetExperience
	ClientboundSetHealth
	ClientboundSetObjective
	ClientboundSetPassengers
	ClientboundSetPlayerTeam
	ClientboundSetScore
	ClientboundSetSimulationDistance
	ClientboundSetSubtitleText
	ClientboundSetTime
	ClientboundSetTitleText
	ClientboundSetTitlesAnimation
	ClientboundSoundEntity
	ClientboundSound
	ClientboundStartConfiguration
	ClientboundStopSound
	ClientboundStoreCookie
	ClientboundSystemChat
	ClientboundTabList
	ClientboundTagQuery
	ClientboundTakeItemEntity
	ClientboundTeleportEntity
	ClientboundTickingState
	ClientboundTickingStep
	ClientboundTransfer
	ClientboundUpdateAdvancements
	ClientboundUpdateAttributes
	ClientboundUpdateMobEffect
	ClientboundUpdateRecipes
	ClientboundUpdateTags
	ClientboundProjectilePower
	ClientboundCustomReportDetails
	ClientboundServerLinks
	ClientboundPacketIDGuard
)

// Game Serverbound
const (
	ServerboundAcceptTeleportation ServerboundPacketID = iota
	ServerboundBlockEntityTagQuery
	ServerboundChangeDifficulty
	ServerboundChatAck
	ServerboundChatCommand
	ServerboundChatCommandSigned
	ServerboundChat
	ServerboundChatSessionUpdate
	ServerboundChunkBatchReceived
	ServerboundClientCommand
	ServerboundClientInformation
	ServerboundCommandSuggestion
	ServerboundConfigurationAcknowledged
	ServerboundContainerButtonClick
	ServerboundContainerClick
	ServerboundContainerClose
	ServerboundContainerSlotStateChanged
	ServerboundCookieResponse
	ServerboundCustomPayload
	ServerboundDebugSampleSubscription
	ServerboundEditBook
	ServerboundEntityTagQuery
	ServerboundInteract
	ServerboundJigsawGenerate
	ServerboundKeepAlive
	ServerboundLockDifficulty
	ServerboundMovePlayerPos
	ServerboundMovePlayerPosRot
	ServerboundMovePlayerRot
	ServerboundMovePlayerStatusOnly
	ServerboundMoveVehicle
	ServerboundPaddleBoat
	ServerboundPickItem
	ServerboundPingRequest
	ServerboundPlaceRecipe
	ServerboundPlayerAbilities
	ServerboundPlayerAction
	ServerboundPlayerCommand
	ServerboundPlayerInput
	ServerboundPong
	ServerboundRecipeBookChangeSettings
	ServerboundRecipeBookSeenRecipe
	ServerboundRenameItem
	ServerboundResourcePack
	ServerboundSeenAdvancements
	ServerboundSelectTrade
	ServerboundSetBeacon
	ServerboundSetCarriedItem
	ServerboundSetCommandBlock
	ServerboundSetCommandMinecart
	ServerboundSetCreativeModeSlot
	ServerboundSetJigsawBlock
	ServerboundSetStructureBlock
	ServerboundSignUpdate
	ServerboundSwing
	ServerboundTeleportToEntity
	ServerboundUseItemOn
	ServerboundUseItem
	ServerboundPacketIDGuard
)