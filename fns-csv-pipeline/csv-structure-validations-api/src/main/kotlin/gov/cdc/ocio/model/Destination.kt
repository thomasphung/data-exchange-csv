package gov.cdc.ocio.model

import com.fasterxml.jackson.annotation.JsonIgnoreProperties
import com.fasterxml.jackson.annotation.JsonProperty

@JsonIgnoreProperties(ignoreUnknown = true)
class Destination {
    @JsonProperty("destination_id")
    var destinationId: String? = null
}