package gov.cdc.ocio.functions

import com.azure.storage.blob.BlobClient
import com.fasterxml.jackson.databind.ObjectMapper
import com.microsoft.azure.functions.ExecutionContext
import com.microsoft.azure.functions.HttpRequestMessage
import com.microsoft.azure.functions.HttpResponseMessage
import com.microsoft.azure.functions.HttpStatus
import gov.cdc.model.Destination
import gov.cdc.utils.Blob
import java.io.IOException
import java.util.*
import java.util.stream.Collectors
import mu.KotlinLogging

class ValidateMultiPartFileUpload {
    fun run(request: HttpRequestMessage<Optional<String>>, context: ExecutionContext, blobClient: BlobClient): HttpResponseMessage {

        val logger = KotlinLogging.logger {}

        var destinations: Array<Destination> = emptyArray()

        try {
            val fileBytes = Blob.toByteArray(blobClient)

            if (fileBytes.isNotEmpty()) {
                val mapper = ObjectMapper()
                destinations = mapper.readValue(
                    fileBytes,
                    Array<Destination>::class.java
                )
            }
        } catch (e: IOException) {
            logger.error(e.message)
            return request.createResponseBuilder(HttpStatus.INTERNAL_SERVER_ERROR).build()
        }

        val destStrs = Arrays.stream(destinations).map { destination: Destination -> destination.destinationId }
            .collect(Collectors.toList())

        return request.createResponseBuilder(HttpStatus.OK).header("Content-Type", "application/json").body(destStrs)
            .build()
    }
}